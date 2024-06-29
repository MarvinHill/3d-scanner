package main

import (
	"github.com/MarvinHill/3d-scanner/internal"
)

func main() {
	updatesChannel := make(chan string)
	scanner := internal.NewScannerDriver(updatesChannel)
	webserver := internal.NewWebserver(scanner, updatesChannel)

	go webserver.Run()
	scanner.Run()
}
