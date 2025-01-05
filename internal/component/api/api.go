/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/uuid"
	"github.com/archnum/sdk.http/api"

	"github.com/archnum/gortoz/internal/component/api/admin"
	"github.com/archnum/gortoz/internal/component/api/ui"
	v1 "github.com/archnum/gortoz/internal/component/api/v1"
)

type (
	implHandler struct {
		api.Manager
		admin *admin.API
		ui    *ui.API
		v1    *v1.API
	}
)

func newHandler(c container.Container) (*implHandler, error) {
	logger := logger.Value(c)

	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	logger = logger.New(id, "api")
	logger.Register()

	p := &api.Params{
		Logger: logger,
	}

	manager := api.New(p)

	admin, err := admin.New(c, manager)
	if err != nil {
		return nil, err
	}

	ui, err := ui.New(c, manager)
	if err != nil {
		return nil, err
	}

	v1, err := v1.New(c, manager)
	if err != nil {
		return nil, err
	}

	impl := &implHandler{
		Manager: manager,
		admin:   admin,
		ui:      ui,
		v1:      v1,
	}

	return impl, nil
}

/*
####### END ############################################################################################################
*/
