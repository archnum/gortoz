/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package v1

import (
	"errors"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.http/api"
	"github.com/archnum/sdk.http/api/apierr"
	"github.com/archnum/sdk.http/api/bind"
	"github.com/archnum/sdk.http/api/render"
)

func (api *API) disableTask(rr render.Renderer) error {
	name, err := bind.PathString(rr, "name")
	if err != nil {
		return err
	}

	if err := api.scheduler.DisableTask(name); err != nil {
		if errors.Is(err, failure.NotFound) {
			return apierr.NotFound(err) ////////////////////////////////////////////////////////////////////////////////
		}

		return err
	}

	return nil
}

func (api *API) enableTask(rr render.Renderer) error {
	name, err := bind.PathString(rr, "name")
	if err != nil {
		return err
	}

	return api.scheduler.EnableTask(name)
}

func (api *API) fireTask(rr render.Renderer) error {
	name, err := bind.PathString(rr, "name")
	if err != nil {
		return err
	}

	return api.scheduler.FireTask(name)
}

func (api *API) tasks(router api.Router) {
	router.Patch("/:name/disable", api.disableTask)
	router.Patch("/:name/enable", api.enableTask)
	router.Put("/:name/fire", api.fireTask)
}

/*
####### END ############################################################################################################
*/
