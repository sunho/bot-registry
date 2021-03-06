package xormstore

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store/xormstore/migrations"
)

type XormStore struct {
	e *xorm.Engine
}

func New(e *xorm.Engine) *XormStore {
	e.SetMapper(core.GonicMapper{})
	e.ShowSQL(false)
	return &XormStore{e}
}

func (x *XormStore) Migrate() error {
	err := migrations.Migrate(x.e)
	if err != nil {
		return err
	}

	err = x.e.Sync(
		new(model.User),
		new(model.UserInvite),
		new(model.UserBot),
		new(model.Bot),
		new(model.BotConfig),
		new(model.BotVolume),
		new(model.BotEnv),
		new(model.Build),
		new(model.BuildLog),
		new(model.Run),
		new(model.RunLog),
	)
	if err != nil {
		return err
	}

	return nil
}
