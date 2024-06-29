package main

import (
	"github.com/MarvinHill/3d-scanner/internal"
)

func main() {

	//webserver := internal.Webserver{}
	scanner := internal.NewScannerDriver()

	scanner.Run()
	//go webserver.Run()

	//runtime.Goexit()
}
