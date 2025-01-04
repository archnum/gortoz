/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"github.com/archnum/sdk.application/container"
)

type (
	Config struct {
		Config map[string]any `ms:"config"`
		Loader string         `ms:"loader"`
	}

	configProvider interface {
		ConfigScheduler() *Config
	}
)

func config(c container.Container) *Config {
	return container.Value[configProvider](c, "config").ConfigScheduler()
}

/*
####### END ############################################################################################################
*/
