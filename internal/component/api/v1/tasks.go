/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package v1

import (
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/bind"
	"github.com/archnum/sdk.http/api/render"
)

func (api *API) disableTask(rr render.Renderer) error {
	name, err := bind.PathString(rr, "name")
	if err != nil {
		return err
	}

	return api.scheduler.DisableTask(name)
}

func (api *API) enableTask(rr render.Renderer) error {
	name, err := bind.PathString(rr, "name")
	if err != nil {
		return err
	}

	return api.scheduler.EnableTask(name)
}

func (api *API) tasks(router api.Router) {
	router.Put("/:name/disable", api.disableTask)
	router.Put("/:name/enable", api.enableTask)
}

/*
####### END ############################################################################################################
*/
