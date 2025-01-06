/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package base

import (
	"sync"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/tracer"
)

type (
	Base struct {
		Executor string `yaml:"executor"`
		Schedule string `yaml:"schedule"`
		Retries  uint   `yaml:"retries"`
		Disabled bool   `yaml:"disabled"`
	}
)

func (b *Base) Clone(disabled bool) *Base {
	return &Base{
		Executor: b.Executor,
		Schedule: b.Schedule,
		Retries:  b.Retries,
		Disabled: disabled,
	}
}

type (
	Wrapper struct {
		name  string
		base  *Base
		mutex sync.Mutex
	}
)

func NewWrapper(name string, base *Base) *Wrapper {
	return &Wrapper{
		name: name,
		base: base,
	}
}

func (w *Wrapper) Name() string {
	return w.name
}

func (w *Wrapper) Disabled() bool {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	return w.base.Disabled
}

func (w *Wrapper) Toggle() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.base.Disabled = !w.base.Disabled
}

func (w *Wrapper) Schedule() string {
	return w.base.Schedule
}

func (w *Wrapper) Run() error {
	tracer.Log("Run task", kv.String("name", w.name)) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	return nil
}

/*
####### END ############################################################################################################
*/
