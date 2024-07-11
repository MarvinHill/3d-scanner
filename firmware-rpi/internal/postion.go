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
	fmt.Println("Position: camera:", position.CameraAxis, "table:", position.TableAxis)
	cameraAxis := position.CameraAxis + movement.CameraAxis
	tableAxis := (position.TableAxis + movement.TableAxis) % MaxRotationOnTableAxis

	return NewPosition(cameraAxis, tableAxis)
}

func (p1 *Position) Equals(p2 Position) bool {
	return p1.CameraAxis == p2.CameraAxis && p1.TableAxis == p2.TableAxis
}

type Position struct {
	CameraAxis int
	TableAxis  int
}

func (pos *Position) Print() {
	fmt.Println("Postion CameraAxis: ", pos.CameraAxis, " TableAxis: ", pos.TableAxis)
}
