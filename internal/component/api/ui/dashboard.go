/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package ui

import (
	"fmt"

	"github.com/archnum/sdk.http/api/render"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/archnum/gortoz/internal/task"
)

func (api *API) disableEnable(task *task.State) g.Node {
	if task.Disabled {
		return A(
			Class("tag is-success"),
			Title("Activer l'exécution de la tâche"),
			g.Attr("hx-patch", fmt.Sprintf("/api/v1/tasks/%s/enable", task.Name)),
			g.Text("activer"),
		)
	}

	return A(
		Class("tag is-danger"),
		Title("Désactiver l'exécution de la tâche"),
		g.Attr("hx-patch", fmt.Sprintf("/api/v1/tasks/%s/disable", task.Name)),
		g.Text("désactiver"),
	)
}

func (api *API) dashboardData(rr render.Renderer) error {
	n := 0
	tasks := api.backend.Tasks()

	return Div(
		Class("table-container"),
		Table(
			Class("table is-bordered is-striped is-narrow is-hoverable is-fullwidth"),
			THead(
				Tr(
					Th(),
					Th(g.Text("Nom")),
					Th(g.Text("Type")),
					Th(g.Text("Planification")),
					Th(g.Text("Le dernier succès")),
					Th(g.Text("Le dernier échec")),
					Th(g.Text("Action")),
					Th(g.Text("Prochaine exécution")),
				),
			),
			TBody(
				g.Group(
					g.Map(
						tasks,
						func(task *task.State) g.Node {
							n++

							return Tr(
								Th(g.Textf("%d", n)),
								Td(
									A(
										g.Attr("hx-put", fmt.Sprintf("/api/v1/tasks/%s/fire", task.Name)),
										Title("Effectuer une exécution immédiate"),
										g.Text(task.Name),
									),
								),
								Td(g.Text(task.Executor)),
								Td(g.Text(task.Schedule)),
								Td(
									g.Text(task.LastSuccess),
									g.Raw("&nbsp;"),
									Span(
										Class("tag is-success"),
										Title("Nombre d'exécutions ayant réussi"),
										g.Textf("%d", task.SuccessCount),
									),
								),
								Td(
									Title(task.ErrMsg),
									g.Text(task.LastFailure),
									g.Raw("&nbsp;"),
									Span(
										Class("tag is-danger"),
										Title("Nombre d'exécutions ayant échoué"),
										g.Textf("%d", task.FailureCount),
									),
								),
								Td(api.disableEnable(task)),
								Td(g.Text(task.NextRun)),
							)
						},
					),
				),
			),
		),
	).Render(rr.ResponseWriter())
}

func (api *API) dashboard(rr render.Renderer) error {
	return api.renderPage(
		rr,
		func() g.Node {
			return g.Group(
				[]g.Node{
					P(
						Class("title is-size-4"),
						g.Text("Tableau de bord"),
					),
					P(
						Class("subtitle is-size-6"),
						g.Text("Liste des tâches planifiées ou désactivées"),
					),
					Div(
						ID("data"),
						g.Attr("hx-trigger", "load, every 10s"),
						g.Attr("hx-get", "/dashboard/data"),
					),
				},
			)
		},
	)
}

/*
####### END ############################################################################################################
*/
