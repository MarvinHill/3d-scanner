package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Webserver struct {
	scannerController *ScannerController
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ws *Webserver) Run() {
	http.HandleFunc("/", handleWebsocket)
	http.ListenAndServe(":8080", nil)
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	log.Println("Websocket connection started")
	conn, _ := upgrader.Upgrade(w, r, nil)
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
			break
		case "PhotoRequest":
			// implement
			break
		}
	}
}

func NewWebserver(scannerController *ScannerController) *Webserver {
	w := &Webserver{}
	w.scannerController = scannerController
	return w
}
