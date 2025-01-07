/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package task

import (
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"

	"github.com/archnum/gortoz/internal/task/fake"
	"github.com/archnum/gortoz/internal/task/http"
)

type (
	Task interface {
		Name() string
		Disabled() bool
		Toggle()
		Executor() string
		Schedule() string
		Run() error
	}
)

func Build(cfg map[string]*Config) (map[string]Task, error) {
	var (
		err  error
		task Task
	)

	tasks := make(map[string]Task)

	for name, tCfg := range cfg {
		switch tCfg.Executor {
		case "fake":
			task, err = fake.NewTask(name, tCfg.Base, tCfg.Config)
		case "http":
			task, err = http.NewTask(name, tCfg.Base, tCfg.Config)
		default:
			return nil,
				failure.New("unknown task executor", kv.String("name", tCfg.Executor)) /////////////////////////////////
		}

		if err != nil {
			return nil, err
		}

		if _, ok := tasks[name]; ok {
			return nil,
				failure.New("this task name is used more than once", kv.String("name", name)) //////////////////////////
		}

		tasks[name] = task
	}

	return tasks, nil
}

/*
####### END ############################################################################################################
*/
