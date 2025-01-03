/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package task

type (
	Config struct {
		Config   map[string]any `ms:"config"`
		Executor string         `ms:"executor"`
		Schedule string         `ms:"schedule"`
		Retries  uint           `ms:"retries"`
		Disabled bool           `ms:"disabled"`
	}
)

/*
####### END ############################################################################################################
*/
