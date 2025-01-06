/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package http

import (
	"github.com/archnum/sdk.base/mapstruct"

	"github.com/archnum/gortoz/internal/task/base"
)

type (
	implTask struct {
		*base.Wrapper
		*config
	}
)

func NewTask(name string, bb *base.Base, cfg map[string]any) (*implTask, error) {
	config := new(config)

	if err := mapstruct.Decode(&config, cfg); err != nil {
		return nil, err
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	impl := &implTask{
		Wrapper: base.NewWrapper(name, bb),
		config:  config,
	}

	return impl, nil
}

/*
####### END ############################################################################################################
*/
