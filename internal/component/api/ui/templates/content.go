/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package templates

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func MainContentTitle(title, subtitle string) g.Node {
	return g.Group(
		[]g.Node{
			P(
				Class("title is-size-4"),
				g.Text(title),
			),
			P(
				Class("subtitle is-size-6"),
				g.Text(subtitle),
			),
		},
	)
}

func MainContentData(link, delay string) g.Node {
	return Div(
		ID("data"),
		g.Attr("hx-trigger", fmt.Sprintf("load, every %s", delay)),
		g.Attr("hx-get", link),
	)
}

func MainContent(title, subtitle, link, delay string) g.Node {
	return g.Group(
		[]g.Node{
			MainContentTitle(title, subtitle),
			MainContentData(link, delay),
		},
	)
}

/*
####### END ############################################################################################################
*/
