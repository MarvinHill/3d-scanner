package internal

func NewPosition(cameraAxis int, tableAxis int) *Position {
	inputCameraAxis := cameraAxis
	inputTableAxis := tableAxis % 360

	if inputCameraAxis < 0 {
		inputCameraAxis = 0
	}
	if inputCameraAxis > 90 {
		inputCameraAxis = 90
	}

	return &Position{CameraAxis: inputCameraAxis, TableAxis: inputTableAxis}
}

type Position struct {
	CameraAxis int
	TableAxis  int
}
