/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package loader

import (
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"

	"github.com/archnum/gortoz/internal/task"
)

func LoadTasks(loader string, cfg map[string]any) (map[string]task.Task, error) {

	switch loader {
	default:
		return nil,
			failure.New("unknown task loader", kv.String("name", loader)) //////////////////////////////////////////////
	}
}

/*
####### END ############################################################################################################
*/
