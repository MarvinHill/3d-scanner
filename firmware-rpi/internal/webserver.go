package internal

import (
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
	http.HandleFunc("/cameraAxisPlus", ws.cameraAxisPlus)
	http.HandleFunc("/cameraAxisMinus", ws.cameraAxisMinus)
	http.HandleFunc("/tableAxisPlus", ws.tableAxisPlus)
	http.HandleFunc("/tableAxisMinus", ws.tableAxisMinus)
	http.HandleFunc("/levelScanner", ws.levelScanner)
	//http.HandleFunc("/takePhoto", nil)
	http.ListenAndServe(":8082", nil)
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
	ws.sc.LevelSites()
}