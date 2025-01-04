/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package main

import (
	"github.com/archnum/sdk.application/component/config"
	"github.com/archnum/sdk.application/component/crypto"
	"github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.application/component/waitend"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/application"
	"github.com/archnum/sdk.http/component/server"

	"github.com/archnum/gortoz/internal/component/cmdline"
	"github.com/archnum/gortoz/internal/component/webui"
	_cfg "github.com/archnum/gortoz/internal/config"
)

var (
	_version string
	_builtAt string
)

func main() {
	app, err := application.New(
		"gortoz",
		application.WithVersion(_version),
		application.WithBuiltAt(_builtAt),
	)
	if err == nil {
		c := container.New(app)

		c.AddComponents( ///////////////////// Liste ordonnée en fonction des dépendances //////////////////////////////
			crypto.New(c),
			config.New(c, new(_cfg.Config)),
			cmdline.New(c),
			logger.New(c),
			webui.New(c),
			server.New(c),
			waitend.New(c),
		)

		err = c.Run( //////////////////////////////// Liste par ordre d'exécution //////////////////////////////////////
			"http.server",
			"waitend",
		)
	}

	app.Exit(err)
}

/*
####### END ############################################################################################################
*/
