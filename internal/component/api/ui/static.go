/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static
var _staticFS embed.FS

func staticFS() (http.FileSystem, error) {
	fsys, err := fs.Sub(_staticFS, "static")
	if err != nil {
		return nil, err
	}

	return http.FS(fsys), nil
}

/*
####### END ############################################################################################################
*/
