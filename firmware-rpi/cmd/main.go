package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/MarvinHill/3d-scanner/internal"
)

var (
	//go:embed resources
	data embed.FS
)

func handleFiles() http.Handler {

	filesys := fs.FS(data)
	html, _ := fs.Sub(filesys, "resources")

	return http.FileServer(http.FS(html))
}

func main() {
	scanner := internal.NewScannerDriver()
	scannerController := internal.NewJobScheduler(scanner)
	webserver := internal.NewWebserver(scannerController, handleFiles())

	scanner.Run()
	webserver.Run()
}
