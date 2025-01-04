/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package yaml

import (
	"os"

	"github.com/archnum/sdk.base/mapstruct"
	"gopkg.in/yaml.v3"

	"github.com/archnum/gortoz/internal/task"
)

type (
	config struct {
		File string `ms:"file"`
	}

	implBackend struct {
		tasks map[string]*task.Config
		file  string
	}
)

func New(data map[string]any) (*implBackend, error) {
	cfg := new(config)
	if err := mapstruct.Decode(&cfg, data); err != nil {
		return nil, err
	}

	impl := &implBackend{
		file: cfg.File,
	}

	return impl, nil
}

func (impl *implBackend) AmITheLeader() bool {
	return true
}

func (impl *implBackend) LoadTasks() (map[string]*task.Config, error) {
	bs, err := os.ReadFile(impl.file)
	if err != nil {
		return nil, err
	}

	var tasks map[string]*task.Config

	if err := yaml.Unmarshal(bs, &tasks); err != nil {
		return nil, err
	}

	impl.tasks = tasks

	return tasks, nil
}

func (impl *implBackend) Close() error {
	return nil
}

/*
####### END ############################################################################################################
*/
