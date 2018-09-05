package store

import (
	"errors"

	"github.com/sunho/fws/server/model"
)

var (
	ErrNoEntry = errors.New("store: no such entry")
)

type Store interface {
	GetUser(id int) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByNickname(nickname string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error

	GetUserInvite(username string) (*model.UserInvite, error)
	CreateUserInvite(i *model.UserInvite) (*model.UserInvite, error)
	DeleteUserInvite(i *model.UserInvite) error

	GetBot(id int) (*model.Bot, error)
	CreateBot(bot *model.Bot) (*model.Bot, error)
	UpdateBot(bot *model.Bot) error
	DeleteBot(bot *model.Bot) error

	//ListUserBot(id int) ([]*model.Bot, error)
	//CreateUserBot(id int, botid int) error
	//DeleteUserBot(userid int, botid int) error

	//ListBotConfig(id int) ([]*model.Config, error)
	//CreateBotConfig(id int, config *model.Config) (*model.Config, error)
	//UpdateBotConfig(id int, config *model.Config) error
	//DeleteBotConfig(id int, config *model.Config) error

	//ListBotVolume(id int) ([]*model.Volume, error)
	//CreateBotVolume(id int, volume *model.Volume) (*model.Volume, error)
	//UpdateBotVolume(id int, volume *model.Volume) error
	//DeleteBotVolume(id int, volume *model.Volume) error

	//ListBotEnv(id int) ([]*model.Env, error)
	//CreateBotEnv(id int, env *model.Env) (*model.Env, error)
	//UpdateBotEnv(id int, env *model.Env) error
	//DeleteBotEnv(id int, env *model.Env) error

	ListBotBuild(bot int) ([]*model.Build, error)
	CreateBotBuild(build *model.Build) (*model.Build, error)
	DeleteBotBuild(build *model.Build) error

	GetBotBuildLog(bot int, number int) (*model.BuildLog, error)
	CreateBotBuildLog(build *model.BuildLog) (*model.BuildLog, error)

	//GetWebhook(hash string) (*model.Webhook, error)
	//CreateWebhook(hook *model.Webhook) (*model.Webhook, error)
	//DeleteWebhook(hook *model.Webhook) error
}
