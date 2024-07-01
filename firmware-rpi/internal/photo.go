package internal

type Photo struct {
	AngleCameraAxis int    `json:"angleCameraAxis"`
	AngleTableAxis  int    `json:"angleTableAxis"`
	PhotoName       string `json:"photoName"`
	PhotoDataBase64 string `json:"photoDataBase64"`
}

type PhotoRequest struct {
	AngleCameraAxis int `json:"angleCameraAxis"`
	AngleTableAxis  int `json:"angleTableAxis"`
}

func (pr *PhotoRequest) ToPosition() *Position {
	return NewPosition(pr.AngleCameraAxis, pr.AngleTableAxis)
}
