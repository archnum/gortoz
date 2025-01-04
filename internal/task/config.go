/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package task

import "github.com/archnum/gortoz/internal/task/base"

type (
	Config struct {
		Config    map[string]any `yaml:"config"`
		base.Base `yaml:",inline"`
	}
)

/*
####### END ############################################################################################################
*/
