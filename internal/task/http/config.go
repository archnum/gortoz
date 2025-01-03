/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package http

import (
	"net/http"
	"time"

	"github.com/archnum/sdk.base/failure"
)

type (
	config struct {
		Method       string        `ms:"method"` // Required
		URL          string        `ms:"url"`
		Timeout      time.Duration `ms:"timeout"`
		ExpectedCode int           `ms:"expected_code"`
	}
)

func (cfg *config) validate() error {
	if cfg.URL == "" {
		return failure.New("attribut 'url' is required") ///////////////////////////////////////////////////////////////
	}

	if cfg.Method == "" {
		cfg.Method = http.MethodGet
	}

	return nil
}

/*
####### END ############################################################################################################
*/
