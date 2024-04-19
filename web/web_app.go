package web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed static

var Static embed.FS

func WebAppHandler() (http.Handler, error) {
	fs, err := fs.Sub(Static, "static")
	if err != nil {
		return nil, fmt.Errorf("unable to wrap embedded filesystem: %w", err)
	}
	handler := http.FileServer(http.FS(fs))
	return handler, nil
}
