/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package ui

import (
	"github.com/archnum/gortoz/internal/component/backend"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/application"
	"github.com/archnum/sdk.http/api"
)

type (
	API struct {
		api.Manager
		app     *application.Application
		backend backend.Backend
	}
)

func New(c container.Container, manager api.Manager) (*API, error) {
	fs, err := staticFS()
	if err != nil {
		return nil, err
	}

	api := &API{
		Manager: manager,
		app:     c.App(),
		backend: backend.Value(c),
	}

	router := manager.Router()

	router.Static(fs)
	router.Get("/", api.dashboard)
	router.Get("/dashboard", api.dashboard)
	router.Get("/dashboard/data", api.dashboardData)

	return api, nil
}

/*
####### END ############################################################################################################
*/
