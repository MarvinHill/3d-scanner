package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Webserver struct {
	scanner         *ScannerDriver
	websConnections []*websocket.Conn
	updates         chan string
}

var upgrader = websocket.Upgrader{}

func (ws *Webserver) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/conn", ws.handleWebsocket).Schemes("wss")
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
		if mapRequest["messageType"] == "manual" {
			var request ManualControl
			err = json.Unmarshal(message, &request)
			if err != nil {
				log.Println("unmarshal:", err)
				continue
			}
			ws.scanner.MoveByManualControl(&request)
		}
	}
}

func NewWebserver(scanner *ScannerDriver, updatesChannel chan string) *Webserver {
	w := &Webserver{}
	w.scanner = scanner
	w.updates = updatesChannel
	return w
}
