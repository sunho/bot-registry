package basic

import (
	"errors"
	"io"

	"github.com/sunho/fws/server/model"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	ErrAlreadyExist = errors.New("basic: already exist")
)

type Runner struct {
	cli *kubernetes.Clientset

	Namespace string
}

func NewRunnerFromCluster(namespace string) (*Runner, error) {
	conf, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		return nil, err
	}

	cli, err := kubernetes.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	return &Runner{
		cli:       cli,
		Namespace: namespace,
	}, nil
}

func (r *Runner) CreateBot(bot *model.RunBot) error {
	deploy, err := r.getDeployment(bot)
	if err == nil {
		return ErrAlreadyExist
	}
	if err != ErrNotExists {
		return err
	}

	confs, err := r.getConfigs(bot)
	if err != nil {
		return err
	}
	if len(configs) != 0 {
		return ErrAlreadyExist
	}

	configCli := r.cli.Core().ConfigMaps()
	deployCli := r.cli.Apps().Deployments()

	confs = r.makeConfigs(bot)
	// potential bug
	for _, conf := range confs {
		_, err = configCli.Create(conf)
		if err != nil {
			return err
		}
	}

	deploy = r.makeDeployment(bot)
	_, err = deployCli.Create(deploy)
	if err != nil {
		return err
	}

	return nil
}

func (r *Runner) UpdateBot(bot *model.RunBot) error {
	return nil
}

func (r *Runner) DeleteBot(bot *model.Bot) error {
	return nil
}

func (r *Runner) RunBot(bot *model.Bot) error {
	return nil
}

func (r *Runner) RestartBot(bot *model.Bot) error {
	return nil
}

func (r *Runner) StopBot(bot *model.Bot) error {
	return nil
}

func (r *Runner) BotStatus(bot *model.Bot) (model.RunStatus, error) {
	return 0, nil
}

func (r *Runner) BotLog(bot *model.Bot) ([]byte, error) {
	return nil, nil
}

func (r *Runner) DownloadVolume(volume *model.BotVolume) (io.Reader, error) {

}

func (r *Runner) VolumeUsed(volume *model.BotVolume) (int64, error) {

}
