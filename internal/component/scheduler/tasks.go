/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

import (
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.http/api/apierr"
)

func (impl *implComponent) DisableTask(name string) error {
	return impl.backend.DisableTask(name)
}

func (impl *implComponent) EnableTask(name string) error {
	return impl.backend.EnableTask(name)
}

func (impl *implComponent) FireTask(name string) error {
	job := impl.getJob(name)
	if job == nil {
		return apierr.NotFound(
			failure.New("this task doesn't exist", kv.String("name", name)), ///////////////////////////////////////////
		)
	}

	job.run()

	return nil
}

/*
####### END ############################################################################################################
*/
