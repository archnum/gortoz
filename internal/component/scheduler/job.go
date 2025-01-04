/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"github.com/archnum/sdk.base/tracer"
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

func (job *job) Run() {
	if !job.manager.AmITheLeader() || job.task.Attr().Disabled {
		return
	}

	tracer.Log(job.task.Name()) // AFAC
}

/*
####### END ############################################################################################################
*/
