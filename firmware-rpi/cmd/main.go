package main

import (
	"github.com/MarvinHill/3d-scanner/internal"
)

func main() {
	scanner := internal.NewScannerDriver()
	scannerController := internal.NewJobScheduler(scanner)
	webserver := internal.NewWebserver(scannerController)

	scanner.Run()
	webserver.Run()
}
