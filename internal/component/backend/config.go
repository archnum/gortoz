/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package backend

import (
	"github.com/archnum/sdk.application/container"
)

type (
	Config struct {
		Config map[string]any `ms:"config"`
		Type   string         `ms:"type"`
	}

	configProvider interface {
		ConfigBackend() *Config
	}
)

func config(c container.Container) *Config {
	return container.Value[configProvider](c, "config").ConfigBackend()
}

/*
####### END ############################################################################################################
*/
