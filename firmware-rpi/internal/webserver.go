package internal

import (
	"encoding/json"
	"fmt"
	"io"
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
	http.HandleFunc("/scanner/takePhoto", ws.takePhoto)
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
	ws.sc.LevelAll()
}

func (ws *Webserver) setScannerLevel(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Webserver request set scanner level")
	ws.sc.SetScannerLevel()
}

func (ws *Webserver) takePhoto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Webserver request take photo")

	data, err := io.ReadAll(r.Body)

	if err != nil || data == nil || len(data) == 0 {
		fmt.Println("Error reading request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !json.Valid(data) {
		fmt.Println("Invalid json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var photoRequest PhotoRequest

	err = json.Unmarshal(data, &photoRequest)

	if err != nil {
		fmt.Println("Error unmarshaling json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo := ws.sc.TakePhoto(photoRequest)
	photoJson, err := json.Marshal(photo)

	if err != nil {
		fmt.Println("Error marshaling photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(photoJson)
}
