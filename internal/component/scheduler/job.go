/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"github.com/robfig/cron/v3"

	"github.com/archnum/gortoz/internal/task"
)

type (
	manager interface {
		AmITheLeader() bool
	}

	job struct {
		task     task.Task
		schedule cron.Schedule
		manager  manager
		entryID  cron.EntryID
	}
)

func (job *job) run() {
	if err := job.task.Run(); err != nil { /////////////////////////////////////////////// Exécution de la tâche ///////
		_ = 0
	}
}

func (job *job) Run() {
	if job.manager.AmITheLeader() {
		if job.task.Disabled() {
			_ = 0
		} else {
			job.run()
		}
	}
}

/*
####### END ############################################################################################################
*/
