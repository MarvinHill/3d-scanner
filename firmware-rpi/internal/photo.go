package internal

type PhotoJob struct {
	JobId                 int     `json:"jobId"`
	PhotoAmountTableAxis  int     `json:"photoAmountTableAxis"`
	PhotoAmountCameraAxis int     `json:"photoAmountCameraAxis"`
	Photo                 []Photo `json:"photo"`
}

type Photo struct {
	AngleCameraAxis int    `json:"angleCameraAxis"`
	AngleTableAxis  int    `json:"angleTableAxis"`
	PhotoName       string `json:"photoName"`
	PhotoDataBase64 string `json:"photoDataBase64"`
}
