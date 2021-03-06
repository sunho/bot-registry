package runtime

import (
	"errors"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

var (
	ErrAlreadyBuilding = errors.New("runtime: already building")
	ErrNotExists       = errors.New("runtime: doesn't exist")
)

const maxCurrent = 10

func NewBuildManager(stor store.Store, builder Builder, runManager *RunManager) *BuildManager {
	return &BuildManager{
		stor:       stor,
		builder:    builder,
		check:      make(chan struct{}),
		builds:     make(map[int]*build),
		runManager: runManager,
	}
}

type BuildManager struct {
	mu sync.RWMutex

	runManager *RunManager
	stor       store.Store
	builder    Builder
	check      chan struct{}
	current    int
	builds     map[int]*build
}

type build struct {
	bot      *model.Bot
	building Building
}

func (b *build) running() bool {
	return b.building != nil
}

func (b *BuildManager) Request(bot *model.Bot) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.builds[bot.ID]; ok {
		return ErrAlreadyBuilding
	}

	b.builds[bot.ID] = &build{
		bot: bot,
	}
	b.check <- struct{}{}

	return nil
}

func (b *BuildManager) Start() {
	go func() {
		for {
			select {
			case <-b.check:
				b.startPendingBuilds()
			}
		}
	}()
}

func (b *BuildManager) Abort(bot *model.Bot) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	bui, ok := b.builds[bot.ID]
	if !ok {
		return ErrNotExists
	}

	if bui.running() {
		return bui.building.Stop()
	}

	delete(b.builds, bui.bot.ID)
	return nil
}

func (b *BuildManager) Status(bot *model.Bot) (model.BuildStatus, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	bui, ok := b.builds[bot.ID]
	if !ok {
		histories, err := b.stor.ListBotBuild(bot.ID)
		if err != nil {
			return nil, err
		}

		recent := new(model.Build)
		for _, history := range histories {
			if recent.Number < history.Number {
				recent = history
			}
		}

		if recent.BotID == 0 {
			return nil, ErrNotExists
		}

		return &model.BuildStatusBuilt{
			Type:    "built",
			Success: recent.Success,
			Number:  recent.Number,
			Created: recent.Created,
		}, nil
	}

	return &model.BuildStatusBuilding{
		Type:    "building",
		Pending: !bui.running(),
		Step:    bui.building.Step(),
	}, nil
}

func (b *BuildManager) startPendingBuilds() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for i := range b.builds {
		if b.current >= maxCurrent {
			return
		}
		bui := b.builds[i]
		if !bui.running() {
			building, err := b.builder.Build(bui.bot, b.callback(bui.bot))
			if err != nil {
				glog.Errorf("Builder.Build faild, err: %v", err)
				continue
			}
			b.current++
			bui.building = building
		}
	}
}

func (b *BuildManager) callback(bot *model.Bot) BuildCallback {
	return func(err error, result string, logged []byte) {
		b.mu.Lock()
		b.current--
		delete(b.builds, bot.ID)
		b.mu.Unlock()
		success := err == nil

		newbot, err := b.stor.GetBot(bot.ID)
		if err != nil {
			glog.Errorf("Error getting bot, err: %v", err)
			return
		}
		newbot.BuildResult = result
		err = b.stor.UpdateBot(newbot)
		if err != nil {
			glog.Errorf("Error updating bot, err: %v", err)
			return
		}

		build, err := b.stor.CreateBotBuild(&model.Build{
			BotID:   bot.ID,
			Success: success,
			Created: time.Now(),
		})
		if err != nil {
			glog.Errorf("Creating Build failed, err: %v", err)
			return
		}

		_, err = b.stor.CreateBotBuildLog(&model.BuildLog{
			BotID:  bot.ID,
			Number: build.Number,
			Logged: logged,
		})
		if err != nil {
			glog.Errorf("Creating BuildLog failed, err: %v", err)
			return
		}

		err = b.runManager.runner.UpdateBuild(newbot)
		if err != nil {
			glog.Errorf("Upading Build failed, err: %v", err)
			return
		}
	}
}
