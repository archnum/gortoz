/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"time"

	"github.com/robfig/cron/v3"

	"github.com/archnum/gortoz/internal/task"
)

type (
	manager interface {
		AmITheLeader() bool
		CollectResult(result *task.Result)
	}

	job struct {
		task     task.Task
		schedule cron.Schedule
		manager  manager
		entryID  cron.EntryID
	}
)

func (job *job) run() {
	now := time.Now()

	result := &task.Result{
		Name:      job.task.Name(),
		Timestamp: now.Format(time.DateTime),
		Disabled:  !job.task.Enabled(),
		Schedule:  job.task.Attr().Schedule,
		Success:   true,
		NextRun:   job.schedule.Next(now).Format(time.DateTime),
	}

	if err := job.task.Run(); err != nil { /////////////////////////////////////////////// Exécution de la tâche ///////
		result.Success = false
		result.Error = err
	}

	job.manager.CollectResult(result)
}

func (job *job) Run() {
	if job.manager.AmITheLeader() {
		if job.task.Enabled() {
			job.run()
		} else {
			result := &task.Result{
				Name:      job.task.Name(),
				Timestamp: time.Now().Format(time.DateTime),
				Disabled:  true,
				Schedule:  job.task.Attr().Schedule,
			}

			job.manager.CollectResult(result)
		}
	}
}

/*
####### END ############################################################################################################
*/
