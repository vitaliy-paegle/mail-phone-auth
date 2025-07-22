package static

import (
	"embed"
	"net/http"
)

//go:embed static/*
var staticFS embed.FS
var fs = http.FS(staticFS)

func New() *http.FileSystem {
	return &fs
}
