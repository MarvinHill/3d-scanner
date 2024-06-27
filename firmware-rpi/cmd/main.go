package main

import (
	"github.com/MarvinHill/3d-scanner/internal"
	"runtime"
)

func main() {

	webserver := internal.Webserver{}
	scanner := internal.ScannerDriver{}

	go scanner.Run()
	go webserver.Run()

	runtime.Goexit()
}
