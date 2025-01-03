/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"github.com/archnum/sdk.application/container"

	"github.com/archnum/gortoz/internal/component/scheduler/loader"
	"github.com/archnum/gortoz/internal/task"
)

const (
	_name = "scheduler"
)

type (
	implComponent struct {
		*container.Component
		tasks map[string]task.Task
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent(_name, c),
		tasks:     make(map[string]task.Task),
	}
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Build() error {
	c := impl.C()
	cfg := config(c)

	_, err := loader.LoadTasks(cfg.Tasks, cfg.Config)
	if err != nil {
		return err
	}

	return nil
}

/*
####### END ############################################################################################################
*/
