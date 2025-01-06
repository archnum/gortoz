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
		RunResult(task task.Task, result *task.Result)
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

	var nextRun string

	if !job.task.Disabled() {
		nextRun = job.schedule.Next(now).Format(time.DateTime)
	}

	result := &task.Result{
		DateTime: now.Format(time.DateTime),
		Success:  true,
		NextRun:  nextRun,
	}

	if err := job.task.Run(); err != nil {
		result.Success = false
		result.Error = err
	}

	job.manager.RunResult(job.task, result) ////////////////////////////////////////////////////////// RunResult ///////
}

func (job *job) Run() {
	if job.manager.AmITheLeader() && !job.task.Disabled() {
		job.run()
	}
}

/*
####### END ############################################################################################################
*/
