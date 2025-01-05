/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

func (impl *implComponent) DisableTask(name string) error {
	return impl.backend.DisableTask(name)
}

func (impl *implComponent) EnableTask(name string) error {
	return impl.backend.EnableTask(name)
}

/*
####### END ############################################################################################################
*/
