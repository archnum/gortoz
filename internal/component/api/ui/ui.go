/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package ui

import (
	"net/http"

	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/render"
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

	router.Get(
		"/static/...",
		func(rr render.Renderer) error {
			http.StripPrefix("/static", http.FileServer(fs)).ServeHTTP(rr.ResponseWriter(), rr.Request())
			return nil
		},
	)

	router.Get("/", api.dashboard)

	return api, nil
}

/*
####### END ############################################################################################################
*/
