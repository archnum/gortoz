/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package base

import "sync"

type (
	Base struct {
		Executor string `yaml:"executor"`
		Schedule string `yaml:"schedule"`
		Retries  uint   `yaml:"retries"`
		Disabled bool   `yaml:"disabled"`
		mutex    sync.Mutex
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

func (b *Base) Enabled() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return !b.Disabled
}

func (b *Base) DisableEnable(disabled bool) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.Disabled = disabled
}

/*
####### END ############################################################################################################
*/
