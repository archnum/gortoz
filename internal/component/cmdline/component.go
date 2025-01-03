/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package cmdline

import (
	"github.com/archnum/sdk.application/component"
	_cpt "github.com/archnum/sdk.application/component/cmdline"
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/application"
	"github.com/archnum/sdk.base/cmdline"

	"github.com/archnum/gortoz/internal/config"
)

func New(c container.Container) component.Component {
	return _cpt.New(
		c,
		func() (*cmdline.CmdLine, error) {
			cfg := container.Value[*config.Config](c, "config")

			cl, err := cmdline.New(
				c.App(),
				cmdline.WithConfigLoader(
					func(_ *application.Application, filepath string) error {
						return config.Load(c, cfg, filepath)
					},
				),
			)
			if err != nil {
				return nil, err
			}

			//////////////
			/// Logger ///
			//////////////

			cl.StringVar(
				&cfg.Logger.Level,
				"level",
				"Logger level",
				cmdline.EnvString("level")...,
			)

			//////////////
			/// Server ///
			//////////////

			cl.StringVar(
				&cfg.Server.Addr,
				"addr",
				"Server address",
				cmdline.EnvString("addr")...,
			)

			cl.StringVar(
				&cfg.Server.CertFile,
				"cert-file",
				"Server certificate file",
				cmdline.EnvString("cert_file")...,
			)

			cl.StringVar(
				&cfg.Server.KeyFile,
				"key-file",
				"Server key file",
				cmdline.EnvString("key_file")...,
			)

			cl.StringVar(
				&cfg.Server.CAFile,
				"ca-file",
				"Server certificate authority file",
				cmdline.EnvString("ca_file")...,
			)

			return cl, nil
		},
	)
}

/*
####### END ############################################################################################################
*/
