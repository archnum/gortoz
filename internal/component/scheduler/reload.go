/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/util"

	"github.com/archnum/gortoz/internal/component/scheduler/loader"
)

func (impl *implComponent) reload() {
	defer func() {
		if data := recover(); data != nil {
			impl.logger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"Reload error recovered",
				kv.Any("data", data),
				kv.String("stack", util.Stack(5)),
			)
		}
	}()

	c := impl.C()
	cfg := config(c)

	impl.logger.Info("Reloading in progress...") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	tasks, err := loader.LoadTasks(cfg.Loader, cfg.Config)
	if err != nil {
		impl.logger.Error("Failed to reload tasks", kv.Error(err)) //:::::::::::::::::::::::::::::::::::::::::::::::::::
		return
	}

	for name, task := range tasks {
		attr := task.Attr()

		pJob, ok := impl.jobs[name]
		if ok {
			pAttr := pJob.task.Attr()

			if attr.Schedule == pAttr.Schedule {
				continue
			}

			impl.cron.Remove(pJob.entryID)
		}

		schedule, err := impl.parser.Parse(attr.Schedule)
		if err != nil {
			impl.logger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"Task schedule error",
				kv.String("name", name),
				kv.String("schedule", attr.Schedule),
			)
		}

		job := &job{
			task:     task,
			schedule: schedule,
		}

		job.entryID = impl.cron.Schedule(schedule, job)

		impl.jobs[name] = job

		impl.logger.Notice( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			util.If(ok, "Modified task", "Added task"),
			kv.String("name", name),
			kv.String("executor", attr.Executor),
			kv.String("schedule", attr.Schedule),
		)
	}

	impl.logger.Info("Reloading complete") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

/*
####### END ############################################################################################################
*/
