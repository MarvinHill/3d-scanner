package internal

var MaxRotationOnTableAxis = 360

func NewPosition(cameraAxis int, tableAxis int) *Position {
	inputCameraAxis := cameraAxis
	inputTableAxis := tableAxis % MaxRotationOnTableAxis

	if inputCameraAxis < 0 {
		inputCameraAxis = 0
	}
	if inputCameraAxis > 90 {
		inputCameraAxis = 90
	}

	return &Position{CameraAxis: inputCameraAxis, TableAxis: inputTableAxis}
}

func AddMovementToPosition(position *Position, movement *Position) *Position {
	cameraAxis := position.CameraAxis + movement.CameraAxis
	tableAxis := (position.TableAxis + movement.TableAxis) % MaxRotationOnTableAxis

	return NewPosition(cameraAxis, tableAxis)
}

type Position struct {
	CameraAxis int
	TableAxis  int
}
