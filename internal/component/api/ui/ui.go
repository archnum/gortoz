/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package ui

import (
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.http/api"
)

type (
	API struct {
		api.Manager
	}
)

func New(_ container.Container, manager api.Manager) (*API, error) {
	fs, err := staticFS()
	if err != nil {
		return nil, err
	}

	api := &API{
		Manager: manager,
	}

	router := manager.Router()

	router.Static(fs)
	router.Get("/", api.dashboard)

	return api, nil
}

/*
####### END ############################################################################################################
*/
