package internal

type JobCreationAnswer struct {
	MessageType string `json:"messageType"`
	Accepted    bool   `json:"accepted"`
	JobId       int    `json:"jobId"`
}

type PhotoJobUpdate struct {
	MessageType string  `json:"messageType"`
	JobId       int     `json:"jobId"`
	Photos      []Photo `json:"photos"`
	Status      string  `json:"status"`
}

// Client Requests
type RequestScan struct {
	MessageType           string `json:"messageType"`
	PhotoAmountTableAxis  int    `json:"photoAmountTableAxis"`
	PhotoAmountCameraAxis int    `json:"photoAmountCameraAxis"`
}

type ManualControl struct {
	MessageType string `json:"messageType"`
	MoveType    string `json:"moveType"`
}
