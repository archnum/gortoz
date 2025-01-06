/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package yaml

import (
	"cmp"
	"log"
	"os"
	"slices"

	"github.com/archnum/gortoz/internal/task"
	_task "github.com/archnum/gortoz/internal/task"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
)

func (impl *implBackend) SetState(task task.Task, nextRun string) {
	state := _task.NewState(task, nextRun)

	impl.mutex.Lock()
	defer impl.mutex.Unlock()

	impl.tasksState[task.Name()] = state
}

func (impl *implBackend) RunResult(task task.Task, result *_task.Result) {
	name := task.Name()

	impl.mutex.Lock()
	defer impl.mutex.Unlock()

	state, ok := impl.tasksState[name]
	if !ok {
		log.New(
			os.Stderr,
			"An error occurred",
			log.LstdFlags|log.Llongfile).Print(
			failure.New("strangely, this task doesn't exist", kv.String("name", name)), //::::::::::::::::::::::::::::::
		)

		return
	}

	state.AfterRun(result)
}

func (impl *implBackend) Tasks() []*_task.State {
	var tasks []*_task.State

	impl.mutex.Lock()

	for _, state := range impl.tasksState {
		tasks = append(tasks, state.Clone())
	}

	impl.mutex.Unlock()

	slices.SortFunc(tasks, func(a, b *_task.State) int { return cmp.Compare(a.Name, b.Name) })

	return tasks
}

/*
####### END ############################################################################################################
*/
