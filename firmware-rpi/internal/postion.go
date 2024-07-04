package internal

import "fmt"

var MaxRotationOnTableAxis = 360

func NewPosition(cameraAxis int, tableAxis int) Position {
	inputCameraAxis := cameraAxis
	inputTableAxis := tableAxis % MaxRotationOnTableAxis

	if inputCameraAxis < 0 {
		inputCameraAxis = 0
	}
	if inputCameraAxis > 90 {
		inputCameraAxis = 90
	}

	return Position{CameraAxis: inputCameraAxis, TableAxis: inputTableAxis}
}

func AddMovementToPosition(position Position, movement Position) Position {
	fmt.Println("Adding movement to position")
	cameraAxis := position.CameraAxis + movement.CameraAxis
	tableAxis := (position.TableAxis + movement.TableAxis) % MaxRotationOnTableAxis

	return NewPosition(cameraAxis, tableAxis)
}

type Position struct {
	CameraAxis int
	TableAxis  int
}

func (pos *Position) Print() {
	fmt.Println("Postion CameraAxis: ", pos.CameraAxis, " TableAxis: ", pos.TableAxis)
}
