/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"context"
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
	Scheduler interface{}

	implComponent struct {
		*container.Component
		logger    *_logger.Logger
		cron      *cron.Cron
		jobs      map[string]*job
		goTracker *gotracker.GoTracker
		parser    cron.Parser
		backend   backend.Backend
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

func (impl *implComponent) addJob(task task.Task, schedule cron.Schedule) {
	job := &job{
		task:     task,
		schedule: schedule,
	}

	job.entryID = impl.cron.Schedule(schedule, job)

	impl.jobs[task.Name()] = job
}

//////////////////////
/// Implementation ///
//////////////////////

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
		attr := task.Attr()
		schedule, err := parser.Parse(attr.Schedule)
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
