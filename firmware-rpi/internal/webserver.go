package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Webserver struct {
	scannerController *ScannerController
	websConnections   []*websocket.Conn
	fileHandler       http.Handler
}

var upgrader = websocket.Upgrader{}

func (ws *Webserver) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/conn", ws.handleWebsocket).Schemes("ws")
	r.Handle("/", ws.fileHandler)
	http.ListenAndServe(":8080", r)
}

func (ws *Webserver) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	ws.websConnections = append(ws.websConnections, conn)
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			continue
		}
		var mapRequest map[string]interface{}
		err = json.Unmarshal(message, &mapRequest)
		if err != nil {
			log.Println("unmarshal:", err)
			continue
		}
		switch mapRequest["messageType"] {
		case "ManualControl":
			var request ManualControlMessage
			err = json.Unmarshal(message, &request)
			if err != nil {
				log.Println("unmarshal:", err)
				continue
			}
			ws.scannerController.MoveByManualControl(&request)
			break
		case "PhotoRequest":
			// implement
			break
		}
	}
}

func NewWebserver(scannerController *ScannerController, fileHandler http.Handler) *Webserver {
	w := &Webserver{}
	w.scannerController = scannerController
	w.fileHandler = fileHandler
	return w
}
