package internal

type ScanJobMessage struct {
	MessageType   string         `json:"messageType"`
	PhotoRequests []PhotoRequest `json:"photoRequests"`
}

type ManualControlMessage struct {
	MessageType string `json:"messageType"`
	MoveType    string `json:"moveType"`
}

type StatusUpdateMessage struct {
	MessageType string `json:"messageType"`
	Status      string `json:"status"`
}
