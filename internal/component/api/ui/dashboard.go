/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package ui

import (
	"github.com/archnum/sdk.http/api/render"
	g "github.com/maragudk/gomponents"

	"github.com/archnum/gortoz/internal/component/api/ui/templates"
)

func (api *API) dashboardData(rr render.Renderer) error {
	rr.NoContent()
	return nil
}

func (api *API) dashboard(rr render.Renderer) error {
	return api.renderPage(
		rr,
		func() g.Node {
			return templates.MainContent(
				"Tableau de bord",
				"Liste des tâches planifiées ou désactivées",
				"/dashboard/data",
				"10s",
			)
		},
	)
}

/*
####### END ############################################################################################################
*/
