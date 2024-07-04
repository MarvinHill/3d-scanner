package main

import (
	"fmt"
	"time"

	"github.com/MarvinHill/3d-scanner/internal"
)

func main() {
	scanner := internal.NewScannerDriver()
	webserver := internal.NewWebserver(scanner)

	scanner.Run()
	go func() {
		for {

			fmt.Println("Enter movement: ")
			scanner.MoveByManualControl("tb_pl")
			scanner.MoveByManualControl("c_pl")
			time.Sleep(1 * time.Second)
		}
	}()
	webserver.Run()

}
