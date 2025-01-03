/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/logger/level"
	"github.com/robfig/cron/v3"
)

type (
	cronLogger struct {
		logger *logger.Logger
	}
)

func newCronLogger(logger *logger.Logger) cron.Logger {
	return &cronLogger{logger}
}

func (cl *cronLogger) Info(msg string, kv ...any) {}

func (cl *cronLogger) Error(err error, msg string, args ...any) {
	kvs := append(logger.ArgsToKV(args), kv.Error(err))

	cl.logger.FormatAndLog(level.Error, "[scheduler] "+msg, kvs...) //::::::::::::::::::::::::::::::::::::::::::::::::::
}

/*
####### END ############################################################################################################
*/
