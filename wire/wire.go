//go:build wireinject
// +build wireinject

package wire

import (
	"test-bot/app"
	"test-bot/config"
	"test-bot/logger"
	"test-bot/login"

	"github.com/google/wire"
)

func InitializeApp() (app.App, error) {
	panic(
		wire.Build(
			app.NewTestBotApp,
			logger.NewLogger,
			config.NewConfig,
			login.NewLoginRepository,

		),
	)
}
