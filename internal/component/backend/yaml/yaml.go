/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package yaml

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/mapstruct"
	"github.com/archnum/sdk.base/uuid"
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
		mutex sync.Mutex
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
	impl.mutex.Lock()
	defer impl.mutex.Unlock()

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

func (impl *implBackend) disableEnable(name string, disabled bool) (bool, error) {
	impl.mutex.Lock()
	defer impl.mutex.Unlock()

	task, ok := impl.tasks[name]
	if !ok {
		return false,
			failure.New("this task doesn't exist", kv.String("name", name)) ////////////////////////////////////////////
	}

	if task.Enabled() != disabled {
		return false, nil
	}

	task.DisableEnable(disabled)

	bs, err := yaml.Marshal(impl.tasks)
	if err != nil {
		return true, err
	}

	id, err := uuid.String()
	if err != nil {
		return true, err
	}

	path := filepath.Dir(impl.file)
	tmpFile := filepath.Join(path, id)

	if err := os.WriteFile(tmpFile, bs, 0644); err != nil {
		return true, err
	}

	if err := os.Rename(tmpFile, impl.file); err != nil {
		_ = os.Remove(tmpFile)
		return true, err
	}

	return true, nil
}

func (impl *implBackend) DisableTask(name string) (bool, error) {
	return impl.disableEnable(name, true)
}

func (impl *implBackend) EnableTask(name string) (bool, error) {
	return impl.disableEnable(name, false)
}

func (impl *implBackend) Close() error {
	return nil
}

/*
####### END ############################################################################################################
*/
