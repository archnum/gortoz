/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package v1

import (
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/middleware"
)

type (
	API struct {
		api.Manager
	}
)

func New(c container.Container, manager api.Manager) (*API, error) {
	api := &API{
		Manager: manager,
	}

	manager.Router().Mount("/api/v1", api.v1)

	return api, nil
}

func (api *API) v1(router api.Router) {

	router.Use(
		middleware.Logger(api.Logger()),
		middleware.Recover(api.Logger()),
	)

	router.Mount("/tasks", api.tasks)
}

/*
####### END ############################################################################################################
*/
