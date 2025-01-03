/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package webui

import (
	"net/http"

	"github.com/archnum/sdk.application/component"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.http/component/handler"
)

func New(c container.Container) component.Component {
	return handler.New(
		c,
		func() (http.Handler, error) {
			return newHandler(c)
		},
	)
}

/*
####### END ############################################################################################################
*/
