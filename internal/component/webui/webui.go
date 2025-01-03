/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package webui

import (
	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/uuid"
	"github.com/archnum/sdk.http/api"
)

type (
	implHandler struct {
		api.Manager
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

	impl := &implHandler{
		Manager: api.New(p),
	}

	impl.declareAPI()

	return impl, nil
}

func (impl *implHandler) declareAPI() {
	router := impl.Router()

	router.Get("/", impl.dashboard)

	router.Mount("/admin", impl.admin)
}

/*
####### END ############################################################################################################
*/
