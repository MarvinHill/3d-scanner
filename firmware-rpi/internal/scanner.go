package internal

import (
	"fmt"
	"sync"

	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

type ScannerDriver struct {
	mu                   sync.Mutex
	Adapter              *raspi.Adaptor
	MotorTableDriver     *gpio.StepperDriver
	MotorOneCameraDriver *gpio.StepperDriver
	MotorTwoCameraDriver *gpio.StepperDriver
	manualStepAmount     int
	CurrentPosition      Position
}

func NewScannerDriver() *ScannerDriver {
	s := &ScannerDriver{}
	s.Adapter = raspi.NewAdaptor()
	s.MotorTableDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"3", "5", "7", "11"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorOneCameraDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"13", "15", "19", "21"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorTwoCameraDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"8", "10", "12", "16"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.manualStepAmount = 10
	return s
}

func (s *ScannerDriver) LevelSites() {
	s.mu.Lock()
	defer s.mu.Unlock()

	// todo check if already level
	oneLevel := false
	twoLevel := false

	for oneLevel == false || twoLevel == false {

		// check level of axis motors

		if oneLevel == false {
			oneFirstCheck, _ := s.Adapter.DigitalRead("36")
			if oneFirstCheck == 1 {
				oneLevel = true
			} else {
				// move down
				fmt.Println("Moving camera axis 1")
				s.MotorOneCameraDriver.Move(4)
				oneSecondCheck, _ := s.Adapter.DigitalRead("36")
				if oneSecondCheck == 1 {
					oneLevel = true
				}
			}

		}

		if twoLevel == false {
			twoFirstCheck, _ := s.Adapter.DigitalRead("38")
			if twoFirstCheck == 1 {
				twoLevel = true
			} else {
				// move down
				fmt.Println("Moving camera axis 2")
				s.MotorTwoCameraDriver.Move(4)
				twoSecondCheck, _ := s.Adapter.DigitalRead("38")
				if twoSecondCheck == 1 {
					twoLevel = true
				}
			}
		}

	}
	s.CurrentPosition = NewPosition(0, 0)
}

func (s *ScannerDriver) TakePhoto(request PhotoRequest) Photo {
	s.mu.Lock()
	defer s.mu.Unlock()

	requestedPos := request.ToPosition()

	tableAxisDiff := requestedPos.TableAxis - s.CurrentPosition.TableAxis // invert because of stepper motor mounting direction
	cameraAxisDiff := requestedPos.CameraAxis - s.CurrentPosition.CameraAxis

	fmt.Println("Moving camera axis 1")
	s.MotorOneCameraDriver.MoveDeg(-cameraAxisDiff * 2)
	fmt.Println("Moving camera axis 2")
	s.MotorTwoCameraDriver.MoveDeg(-cameraAxisDiff * 2)
	fmt.Println("Moving table")
	s.MotorTableDriver.MoveDeg(tableAxisDiff)
	s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, NewPosition(cameraAxisDiff, tableAxisDiff))

	// s.takePhoto()

	return Photo{}
}

func (s *ScannerDriver) MoveByManualControl(movement string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println("Moving by manual control")

	switch movement {
	case "c_pl":
		prevPos := s.CurrentPosition
		s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, NewPosition(s.manualStepAmount, 0))
		s.CurrentPosition.Print()
		if prevPos.Equals(s.CurrentPosition) {
			break
		}
		fmt.Println("Moving camera axis 1")
		s.MotorOneCameraDriver.MoveDeg(-s.manualStepAmount * 2)
		fmt.Println("Moving camera axis 2")
		s.MotorTwoCameraDriver.MoveDeg(-s.manualStepAmount * 2)
		break
	case "c_min":
		prevPos := s.CurrentPosition
		s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, NewPosition(-s.manualStepAmount, 0))
		s.CurrentPosition.Print()
		if prevPos.Equals(s.CurrentPosition) {
			break
		}
		fmt.Println("Moving camera axis 1")
		s.MotorOneCameraDriver.MoveDeg(s.manualStepAmount * 2)
		fmt.Println("Moving camera axis 2")
		s.MotorTwoCameraDriver.MoveDeg(s.manualStepAmount * 2)
		break
	case "tb_pl":
		fmt.Println("Moving table")
		s.MotorTableDriver.MoveDeg(s.manualStepAmount)
		s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, NewPosition(0, s.manualStepAmount))
		s.CurrentPosition.Print()
		break
	case "tb_min":
		fmt.Println("Moving table")
		s.MotorTableDriver.MoveDeg(-s.manualStepAmount)
		s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, NewPosition(0, -s.manualStepAmount))
		s.CurrentPosition.Print()
		break
	}
}

func (s *ScannerDriver) Run() {
	fmt.Println("Starting Scanner")
	s.Adapter.Connect()
	s.MotorTableDriver.Start()
	s.MotorOneCameraDriver.Start()
	s.MotorTwoCameraDriver.Start()

	// Reset Table Motor Output Pins
	s.Adapter.DigitalWrite("3", 0)
	s.Adapter.DigitalWrite("5", 0)
	s.Adapter.DigitalWrite("7", 0)
	s.Adapter.DigitalWrite("11", 0)
	// Reset Camera Motor Output Pins
	s.Adapter.DigitalWrite("13", 0)
	s.Adapter.DigitalWrite("15", 0)
	s.Adapter.DigitalWrite("19", 0)
	s.Adapter.DigitalWrite("21", 0)
	s.Adapter.DigitalWrite("8", 0)
	s.Adapter.DigitalWrite("10", 0)
	s.Adapter.DigitalWrite("12", 0)
	s.Adapter.DigitalWrite("16", 0)

	// Sensor 1
	s.Adapter.DigitalWrite("35", 1)
	// Sensor 2
	s.Adapter.DigitalWrite("37", 1)
}
