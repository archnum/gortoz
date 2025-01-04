/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package base

type (
	Base struct {
		Executor string `yaml:"executor"`
		Schedule string `yaml:"schedule"`
		Retries  uint   `yaml:"retries"`
		Disabled bool   `yaml:"disabled"`
	}
)

/*
####### END ############################################################################################################
*/
