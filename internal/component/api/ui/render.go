/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package ui

import (
	"net/http"
	"time"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

const (
	_bulma = "/static/bulma.1.0.2.css"
	_htmx  = "/static/htmx.2.0.4.js"
)

func (api *API) renderPage(w http.ResponseWriter, content func() g.Node) error {
	app := api.app

	return Doctype(
		HTML(
			//Data("theme", "light"),
			Lang("fr"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				TitleEl(g.Text(app.Name())),
				Link(Rel("stylesheet"), Href(_bulma)),
				Link(Rel("stylesheet"), Href("/static/app.css")),
			),
			Body(
				Class("has-navbar-fixed-top"),
				Nav(
					Class("navbar is-primary is-fixed-top"),
					Div(
						Class("container"),
						Div(
							Class("navbar-brand"),
							A(
								Class("navbar-item"),
								Href("/"),
								g.Text(app.Name()),
							),
							A(
								Class("navbar-burger"),
								Data("target", "navbar_menu"),
								Span(Aria("hidden", "true")),
								Span(Aria("hidden", "true")),
								Span(Aria("hidden", "true")),
								Span(Aria("hidden", "true")),
							),
						),
						Div(
							ID("navbar_menu"),
							Class("navbar-menu"),
							Div(
								Class("navbar-start"),
								A(
									Class("navbar-item"),
									Href("/applications"),
									g.Text("applications"),
								),
							),
						),
					),
				),
				Main(
					Class("section"),
					Div(
						Class("container"),
						Div(
							ID("error"),
							Class("notification has-text-danger has-text-centered"),
						),
						content(),
					),
				),
				Footer(
					Class("footer"),
					Div(
						Class("is-size-7 has-text-centered"),
						P(
							g.Textf("%s - v%s - %s", app.Name(), app.Version(), app.BuiltAt().Format(time.DateTime)),
						),
						P(
							g.Raw("archivage numérique"),
						),
					),
				),
				Script(Src(_htmx)),
				Script(Src("/static/app.js")),
			),
		),
	).Render(w)
}

/*
####### END ############################################################################################################
*/
