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
		tasks      map[string]*task.Config
		tasksState map[string]*task.State
		file       string
		mutex      sync.Mutex
	}
)

func New(data map[string]any) (*implBackend, error) {
	cfg := new(config)
	if err := mapstruct.Decode(&cfg, data); err != nil {
		return nil, err
	}

	impl := &implBackend{
		file:       cfg.File,
		tasksState: make(map[string]*task.State),
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

func (impl *implBackend) writeFile() error {
	bs, err := yaml.Marshal(impl.tasks)
	if err != nil {
		return err
	}

	id, err := uuid.String()
	if err != nil {
		return err
	}

	path := filepath.Dir(impl.file)
	tmpFile := filepath.Join(path, id)

	if err := os.WriteFile(tmpFile, bs, 0644); err != nil {
		return err
	}

	if err := os.Rename(tmpFile, impl.file); err != nil {
		_ = os.Remove(tmpFile)
		return err
	}

	return nil
}

func (impl *implBackend) disableEnable(task task.Task, disabled bool) error {
	name := task.Name()

	impl.mutex.Lock()
	defer impl.mutex.Unlock()

	tCfg, ok := impl.tasks[name]
	if !ok {
		return failure.New("strangely, this task doesn't exist", kv.String("name", name)) //////////////////////////////
	}

	if task.Disabled() == disabled {
		return nil
	}

	backup := tCfg.Base

	defer func() {
		impl.tasks[name].Base = backup
	}()

	impl.tasks[name].Base = backup.Clone(disabled)

	if err := impl.writeFile(); err != nil {
		return failure.WithMessage(err, "failed to update file", kv.String("name", impl.file)) /////////////////////////
	}

	task.Toggle()

	state := impl.tasksState[name]
	state.Disabled = disabled
	state.NextRun = ""

	return nil
}

func (impl *implBackend) DisableTask(task task.Task) error {
	return impl.disableEnable(task, true)
}

func (impl *implBackend) EnableTask(task task.Task) error {
	return impl.disableEnable(task, false)
}

func (impl *implBackend) Close() error {
	return nil
}

/*
####### END ############################################################################################################
*/
