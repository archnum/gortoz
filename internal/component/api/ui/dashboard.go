/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package ui

import (
	"github.com/archnum/sdk.http/api/render"
	g "github.com/maragudk/gomponents"
)

func (api *API) dashboard(rr render.Renderer) error {
	return g.Text("Hello world !").Render(rr.ResponseWriter())
}

/*
####### END ############################################################################################################
*/
