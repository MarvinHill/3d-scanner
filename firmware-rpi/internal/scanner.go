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
	// Implement
	s.CurrentPosition = NewPosition(0, 0)
}

func (s *ScannerDriver) TakePhoto(request PhotoRequest) Photo {
	s.mu.Lock()
	defer s.mu.Unlock()

	// s.moveToPostion()
	// s.takePhoto()

	// Implement
	return Photo{}
}

func (s *ScannerDriver) MoveByManualControl(movement string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println("Moving by manual control")

	switch movement {
	case "c_pl":
		fmt.Println("Moving camera axis 1")
		s.MotorOneCameraDriver.MoveDeg(s.manualStepAmount * 2)
		fmt.Println("Moving camera axis 2")
		s.MotorTwoCameraDriver.MoveDeg(-s.manualStepAmount * 2)
		s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, NewPosition(s.manualStepAmount, 0))
		s.CurrentPosition.Print()
		break
	case "c_min":
		fmt.Println("Moving camera axis 1")
		s.MotorOneCameraDriver.MoveDeg(-s.manualStepAmount * 2)
		fmt.Println("Moving camera axis 2")
		s.MotorTwoCameraDriver.MoveDeg(s.manualStepAmount * 2)
		s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, NewPosition(-s.manualStepAmount, 0))
		s.CurrentPosition.Print()
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

}
