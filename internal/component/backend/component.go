/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package backend

import (
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"

	"github.com/archnum/gortoz/internal/component/backend/yaml"
	"github.com/archnum/gortoz/internal/task"
)

const (
	_name = "backend"
)

type (
	Backend interface {
		AmITheLeader() bool
		LoadTasks() (map[string]*task.Config, error)
		DisableTask(task task.Task) error
		EnableTask(task task.Task) error
		SetState(task task.Task, nextRun string)
		RunResult(task task.Task, result *task.Result)
		Tasks() []*task.State
		Close() error
	}

	implComponent struct {
		*container.Component
		backend Backend
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent(_name, c),
	}
}

func Value(c container.Container) Backend {
	return container.Value[Backend](c, _name)
}

/////////////////
/// Component ///
/////////////////

func (impl *implComponent) Build() error {
	c := impl.C()
	cfg := config(c)

	var err error

	switch cfg.Type {
	case "yaml":
		impl.backend, err = yaml.New(cfg.Config)
	default:
		return failure.New("unknown backend type", kv.String("type", cfg.Type)) ////////////////////////////////////////
	}

	if err != nil {
		return err
	}

	impl.SetValue(impl.backend)

	return nil
}

func (impl *implComponent) Close() error {
	return impl.backend.Close()
}

/*
####### END ############################################################################################################
*/
