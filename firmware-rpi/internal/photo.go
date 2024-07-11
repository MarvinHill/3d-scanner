package internal

type Photo struct {
	AngleCameraAxis int    `json:"angleCameraAxis"`
	AngleTableAxis  int    `json:"angleTableAxis"`
	PhotoData       string `json:"photoData"`
}

type PhotoRequest struct {
	AngleCameraAxis int `json:"angleCameraAxis"`
	AngleTableAxis  int `json:"angleTableAxis"`
}

func (pr *PhotoRequest) ToPosition() Position {
	return NewPosition(pr.AngleCameraAxis, pr.AngleTableAxis)
}
