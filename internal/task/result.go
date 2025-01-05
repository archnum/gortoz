/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package task

type (
	// TODO: un pool ?
	Result struct {
		Name      string
		Timestamp string
		Disabled  bool
		Schedule  string
		Success   bool
		Error     error
		NextRun   string
	}
)

/*
####### END ############################################################################################################
*/
