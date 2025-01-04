/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package task

import "github.com/archnum/gortoz/internal/task/base"

type (
	Config struct {
		base.Base `yaml:",inline"`
		Config    map[string]any `yaml:"config"`
	}
)

/*
####### END ############################################################################################################
*/
