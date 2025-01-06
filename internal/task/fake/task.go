/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package fake

import "github.com/archnum/gortoz/internal/task/base"

type (
	implTask struct {
		*base.Wrapper
	}
)

func NewTask(name string, bb *base.Base, _ map[string]any) (*implTask, error) {
	impl := &implTask{
		Wrapper: base.NewWrapper(name, bb),
	}

	return impl, nil
}

/*
####### END ############################################################################################################
*/
