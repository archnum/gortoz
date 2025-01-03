/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/gotracker"
	"github.com/archnum/sdk.base/kv"
	_logger "github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/uuid"
	"github.com/robfig/cron/v3"

	"github.com/archnum/gortoz/internal/component/scheduler/loader"
)

const (
	_name = "scheduler"
)

type (
	implComponent struct {
		*container.Component
		logger    *_logger.Logger
		parser    cron.Parser
		cron      *cron.Cron
		jobs      map[string]*job
		goTracker *gotracker.GoTracker
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent(_name, c),
		jobs:      make(map[string]*job),
	}
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Build() error {
	c := impl.C()
	cfg := config(c)

	id, err := uuid.New()
	if err != nil {
		return err
	}

	impl.logger = logger.Value(c).New(id, "_name")

	logger := newCronLogger(logger.Value(c).New(id, "_name"))
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

	tasks, err := loader.LoadTasks(cfg.Tasks, cfg.Config)
	if err != nil {
		return err
	}

	for name, task := range tasks {
		attr := task.Attr()
		schedule, err := parser.Parse(attr.Schedule)
		if err != nil {
			return failure.WithMessage(err, "task schedule error", kv.String("name", name)) ////////////////////////////
		}

		job := &job{
			task:     task,
			schedule: schedule,
		}

		job.entryID = cron.Schedule(schedule, job)

		impl.jobs[name] = job
	}

	return nil
}

func (impl *implComponent) Start() error {
	impl.goTracker = gotracker.New(gotracker.WithLogger(impl.logger))

	impl.goTracker.Go( //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		_name,
		func(_ chan struct{}) {
			impl.cron.Start()
		},
	)

	return nil
}

func (impl *implComponent) Stop() error {
	<-impl.cron.Stop().Done()

	impl.goTracker.Stop()
	impl.goTracker.Wait()

	return nil
}

/*
####### END ############################################################################################################
*/
