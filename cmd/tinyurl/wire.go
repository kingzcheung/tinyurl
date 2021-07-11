//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kingzcheung/tinyurl/config"
)

func InitializeApp(config *config.Config) (*app, error) {
	wire.Build(
		storeSet,
		serverSet,
		newApp,
	)

	return &app{}, nil
}
