package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Webserver struct {
	sc *ScannerDriver
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ws *Webserver) Run() {
	http.HandleFunc("/scanner/cameraAxisPlus", ws.cameraAxisPlus)
	http.HandleFunc("/scanner/cameraAxisMinus", ws.cameraAxisMinus)
	http.HandleFunc("/scanner/tableAxisPlus", ws.tableAxisPlus)
	http.HandleFunc("/scanner/tableAxisMinus", ws.tableAxisMinus)
	http.HandleFunc("/scanner/levelScanner", ws.levelScanner)
	http.HandleFunc("/scanner/setScannerLevel", ws.setScannerLevel)
	//http.HandleFunc("/scanner/takePhoto", ws.takePhoto)
	http.HandleFunc("/", logAll)
	http.ListenAndServe(":8082", nil)
}

func logAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to: " + r.URL.Path)
}

func NewWebserver(scanner *ScannerDriver) *Webserver {
	w := &Webserver{}
	w.sc = scanner
	return w
}

func (ws *Webserver) cameraAxisPlus(w http.ResponseWriter, r *http.Request) {
	ws.sc.MoveByManualControl("c_pl")
}

func (ws *Webserver) cameraAxisMinus(w http.ResponseWriter, r *http.Request) {
	ws.sc.MoveByManualControl("c_min")
}

func (ws *Webserver) tableAxisPlus(w http.ResponseWriter, r *http.Request) {
	ws.sc.MoveByManualControl("tb_pl")
}

func (ws *Webserver) tableAxisMinus(w http.ResponseWriter, r *http.Request) {
	ws.sc.MoveByManualControl("tb_min")
}

func (ws *Webserver) levelScanner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Webserver request leveling scanner")
	ws.sc.LevelSites()
}

func (ws *Webserver) setScannerLevel(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Webserver request set scanner level")
	ws.sc.SetScannerLevel()
}
