/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package http

import (
	"github.com/archnum/gortoz/internal/task/base"
	"github.com/archnum/sdk.base/tracer"
)

type (
	implTask struct {
		*base.Base
		name string
	}
)

func NewTask(name string, base *base.Base, _ map[string]any) (*implTask, error) {
	impl := &implTask{
		Base: base,
		name: name,
	}

	return impl, nil
}

func (impl *implTask) Name() string {
	return impl.name
}

func (impl *implTask) Attr() *base.Base {
	return impl.Base
}

func (impl *implTask) Run() error {
	tracer.Log(impl.name)
	return nil
}

/*
####### END ############################################################################################################
*/
