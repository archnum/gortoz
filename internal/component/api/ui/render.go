/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package ui

import (
	"fmt"
	"time"

	"github.com/archnum/sdk.http/api/render"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

const (
	_bulma = "/static/bulma.1.0.2.css"
	_htmx  = "/static/htmx.2.0.4.js"
)

func (api *API) renderPage(rr render.Renderer, content func() g.Node) error {
	app := api.app

	return Doctype(
		HTML(
			//Data("theme", "light"),
			Lang("fr"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				Meta(Name("description"), Content(fmt.Sprintf("écosystème %s", app.Ecosystem()))),
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
								Class("navbar-item has-text-weight-semibold"),
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
									g.Text("documentation"),
								),
							),
							Div(
								Class("navbar-end"),
								Div(
									Class("navbar-item has-dropdown is-hoverable"),
									A(
										Class("navbar-link"),
										g.Text("admin"),
									),
									Div(
										Class("navbar-dropdown"),
										A(
											Class("navbar-item"),
											Href("/admin/loggers"),
											g.Text("loggers"),
										),
									),
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
							g.Textf(
								"%s - %s - v%s - %s",
								app.Name(),
								app.ShortDesc(),
								app.Version(),
								app.BuiltAt().Format(time.DateTime),
							),
						),
						P(
							g.Text("écosystème"),
							g.Raw("&nbsp;"),
							Span(
								Class("tag is-black"),
								g.Text(app.Ecosystem()),
							),
						),
						P(g.Text("Copyright (c) 2025 Archivage Numérique")),
					),
				),
				Script(Src(_htmx)),
				Script(Src("/static/app.js")),
			),
		),
	).Render(rr.ResponseWriter())
}

/*
####### END ############################################################################################################
*/
