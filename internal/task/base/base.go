/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package base

type (
	Base struct {
		Executor string `ms:"executor"`
		Schedule string `ms:"schedule"`
		Retries  uint   `ms:"retries"`
		Disabled bool   `ms:"disabled"`
	}
)

/*
####### END ############################################################################################################
*/
