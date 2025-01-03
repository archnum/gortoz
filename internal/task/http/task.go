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
		*base.Base
		*config
		name string
	}
)

func NewTask(name string, base *base.Base, cfg map[string]any) (*implTask, error) {
	config := new(config)

	if err := mapstruct.Decode(&config, cfg); err != nil {
		return nil, err
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	impl := &implTask{
		Base:   base,
		config: config,
		name:   name,
	}

	return impl, nil
}

func (impl *implTask) Name() string {
	return impl.name
}

func (impl *implTask) Attr() *base.Base {
	return impl.Base
}

/*
####### END ############################################################################################################
*/
