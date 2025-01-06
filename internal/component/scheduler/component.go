/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"context"
	"sync"
	"time"

	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/gotracker"
	"github.com/archnum/sdk.base/kv"
	_logger "github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/uuid"
	"github.com/robfig/cron/v3"

	"github.com/archnum/gortoz/internal/component/backend"
	"github.com/archnum/gortoz/internal/task"
)

const (
	Name                = "scheduler"
	_reloadingFrequency = time.Duration(5 * time.Minute)
)

type (
	Scheduler interface {
		DisableTask(name string) error
		EnableTask(name string) error
		FireTask(name string) error
	}

	implComponent struct {
		backend backend.Backend
		*container.Component
		logger    *_logger.Logger
		cron      *cron.Cron
		jobs      map[string]*job
		goTracker *gotracker.GoTracker
		parser    cron.Parser
		mutex     sync.Mutex
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent(Name, c),
		jobs:      make(map[string]*job),
	}
}

func Value(c container.Container) Scheduler {
	return container.Value[Scheduler](c, Name)
}

func (impl *implComponent) getJob(name string) *job {
	impl.mutex.Lock()
	defer impl.mutex.Unlock()

	job, ok := impl.jobs[name]
	if !ok {
		return nil
	}

	return job
}

func (impl *implComponent) addJob(task task.Task, schedule cron.Schedule) {
	job := &job{
		task:     task,
		schedule: schedule,
		manager:  impl.backend,
	}

	job.entryID = impl.cron.Schedule(schedule, job)

	var nextRun string
	disabled := task.Disabled()

	if !disabled {
		nextRun = schedule.Next(time.Now()).Format(time.DateTime)
	}

	impl.backend.SetState(task, nextRun) ////////////////////////////////////////////////////////////// SetState ///////

	impl.mutex.Lock()
	defer impl.mutex.Unlock()

	impl.jobs[task.Name()] = job
}

func (impl *implComponent) deleteJob(job *job) {
	impl.cron.Remove(job.entryID)

	impl.mutex.Lock()
	defer impl.mutex.Unlock()

	delete(impl.jobs, job.task.Name())
}

/////////////////
/// Component ///
/////////////////

func (impl *implComponent) Build() error {
	c := impl.C()

	impl.backend = backend.Value(c)

	tmp, err := impl.backend.LoadTasks()
	if err != nil {
		return err
	}

	tasks, err := task.Build(tmp)
	if err != nil {
		return err
	}

	id, err := uuid.New()
	if err != nil {
		return err
	}

	impl.logger = logger.Value(c).New(id, Name)
	impl.logger.Register()

	logger := newCronLogger(impl.logger)
	parser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	impl.parser = parser

	cron := cron.New(
		cron.WithChain(
			cron.Recover(logger),
			cron.SkipIfStillRunning(logger),
		),
		cron.WithLogger(logger),
		cron.WithParser(parser),
	)

	impl.cron = cron

	for name, task := range tasks {
		schedule, err := parser.Parse(task.Schedule())
		if err != nil {
			return failure.WithMessage(err, "task schedule error", kv.String("name", name)) ////////////////////////////
		}

		impl.addJob(task, schedule)
	}

	impl.SetValue(impl)

	return nil
}

func (impl *implComponent) Start() error {
	impl.goTracker = gotracker.New(gotracker.WithLogger(impl.logger))

	impl.goTracker.Go( //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		Name,
		func(ctx context.Context) error {
			impl.cron.Start()

			for {
				select {
				case <-ctx.Done():
					return nil
				case <-time.After(_reloadingFrequency):
					impl.reload()
				}
			}
		},
	)

	return nil
}

func (impl *implComponent) Stop() error {
	impl.goTracker.Stop()
	<-impl.cron.Stop().Done()
	impl.goTracker.Wait()

	return nil
}

/*
####### END ############################################################################################################
*/
