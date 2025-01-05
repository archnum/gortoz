/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package admin

import (
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/middleware"
	"github.com/archnum/sdk.http/api/render"
	_loggers "github.com/archnum/sdk.loggers"
	_external "github.com/ltrochet/loggers"
)

type (
	API struct {
		api.Manager
	}
)

func New(_ container.Container, manager api.Manager) (*API, error) {
	api := &API{
		Manager: manager,
	}

	manager.Router().Mount("/admin", api.admin)

	return api, nil
}

func (api *API) admin(router api.Router) {
	router.Use(
		middleware.Logger(api.Logger()),
		middleware.Recover(api.Logger()),
	)

	router.Get(
		"/loggers",
		func(rr render.Renderer) error {
			_external.Handler(_loggers.All()).ServeHTTP(rr.ResponseWriter(), rr.Request())
			return nil
		},
	)

	// chi.middleware.Profiler() ?
}

/*
####### END ############################################################################################################
*/
