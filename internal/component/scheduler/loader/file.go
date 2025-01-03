/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package loader

import (
	"github.com/archnum/sdk.base/config"
	"github.com/archnum/sdk.base/mapstruct"

	"github.com/archnum/gortoz/internal/task"
)

type (
	cfgFile struct {
		Path string `ms:"path"`
	}
)

func loadFromFile(cfg map[string]any) (map[string]task.Task, error) {
	file := new(cfgFile)

	if err := mapstruct.Decode(file, cfg); err != nil {
		return nil, err
	}

	var tasks map[string]*task.Config

	if err := config.New().DecodeFile(&tasks, file.Path); err != nil {
		return nil, err
	}

	return task.Build(tasks)
}

/*
####### END ############################################################################################################
*/
