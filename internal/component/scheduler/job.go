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
	job struct {
		task     task.Task
		schedule cron.Schedule
		entryID  cron.EntryID
	}
)

func (job *job) Run() {
	if job.task.Attr().Disabled { // TODO: atomic
		return
	}

	tracer.Log(job.task.Name()) // AFAC
}

/*
####### END ############################################################################################################
*/
