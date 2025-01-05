/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package scheduler

func (impl *implComponent) DisableTask(name string) error {
	_, _ = impl.backend.DisableTask(name)
	return nil
}

func (impl *implComponent) EnableTask(name string) error {
	_, _ = impl.backend.EnableTask(name)
	return nil
}

/*
####### END ############################################################################################################
*/
