/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package config

import (
	"github.com/archnum/gortoz/internal/component/scheduler"
	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/config"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/mapstruct"
	"github.com/archnum/sdk.http/server"
)

type (
	Config struct {
		Logger    logger.Config    `ms:"logger"`
		Scheduler scheduler.Config `ms:"scheduler"`
		Server    server.Config    `ms:"server"`
	}
)

func (cfg *Config) ConfigLogger() *logger.Config {
	return &cfg.Logger
}

func (cfg *Config) ConfigScheduler() *scheduler.Config {
	return &cfg.Scheduler
}

func (cfg *Config) ConfigServer() *server.Config {
	return &cfg.Server
}

func Load(_ container.Container, to *Config, filepath string) error {
	tmp := struct {
		Config map[string]any `ms:"config"`
		Loader string         `ms:"loader"`
	}{}

	err := config.New().DecodeFile(&tmp, filepath)
	if err != nil {
		return err
	}

	switch tmp.Loader {
	case "no", "none":
		err = mapstruct.Decode(to, tmp.Config)
	default:
		err = failure.New("this config loader is unknown", kv.String("name", tmp.Loader)) //////////////////////////////
	}

	return err
}

/*
####### END ############################################################################################################
*/
