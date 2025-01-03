/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package webui

import (
	"github.com/archnum/sdk.http/api/render"
	g "github.com/maragudk/gomponents"
)

func (impl *implHandler) dashboard(rr render.Renderer) error {
	return g.Text("Hello world !").Render(rr.ResponseWriter())
}

/*
####### END ############################################################################################################
*/
