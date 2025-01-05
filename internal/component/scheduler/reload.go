/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/util"

	"github.com/archnum/gortoz/internal/task"
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
	impl.logger.Info("Reloading in progress...") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	tmp, err := impl.backend.LoadTasks()
	if err != nil {
		impl.logger.Error("Failed to reload tasks", kv.Error(err)) //:::::::::::::::::::::::::::::::::::::::::::::::::::
		return
	}

	tasks, err := task.Build(tmp)
	if err != nil {
		impl.logger.Error("Failed to rebuild tasks", kv.Error(err)) //::::::::::::::::::::::::::::::::::::::::::::::::::
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

			impl.deleteJob(pJob)
		}

		schedule, err := impl.parser.Parse(attr.Schedule)
		if err != nil {
			impl.logger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"Task schedule error",
				kv.String("name", name),
				kv.String("schedule", attr.Schedule),
			)

			continue
		}

		impl.addJob(task, schedule)

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
