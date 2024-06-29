package internal

import (
	"fmt"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/raspi"
	"sync"
)

type ScannerDriver struct {
	mu                   sync.Mutex
	level                bool
	Adapter              *raspi.Adaptor
	MotorTableDriver     *gpio.StepperDriver
	MotorOneCameraDriver *gpio.StepperDriver
	MotorTwoCameraDriver *gpio.StepperDriver
	manualStepAmount     int
	CurrentPosition      *Position
	TakePhotoPosition    *Position
	updates              chan string
}

func NewScannerDriver(updates chan string) *ScannerDriver {
	s := &ScannerDriver{}
	s.Adapter = raspi.NewAdaptor()
	s.MotorTableDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"3", "5", "7", "11"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorOneCameraDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"13", "15", "19", "21"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorTwoCameraDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"8", "10", "12", "16"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.TakePhotoPosition = NewPosition(0, 0)
	s.manualStepAmount = 10
	s.updates = updates
	return s
}

func (s *ScannerDriver) LevelSites() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.level = true
	s.CurrentPosition = NewPosition(0, 0)
}

func (s *ScannerDriver) GetLevel() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.level
}

func (s *ScannerDriver) MoveByManualControl(control *ManualControl) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch control.MoveType {
	case "camera_plus":
		s.MotorOneCameraDriver.Move(s.manualStepAmount)
		s.MotorTwoCameraDriver.Move(-s.manualStepAmount)
		break
	case "camera_minus":
		s.MotorOneCameraDriver.Move(-s.manualStepAmount)
		s.MotorTwoCameraDriver.Move(s.manualStepAmount)
		break
	case "table_plus":
		s.MotorTableDriver.Move(s.manualStepAmount)
		break
	case "table_minus":
		s.MotorTableDriver.Move(-s.manualStepAmount)
		break
	}
}

func (s *ScannerDriver) Run() {
	fmt.Println("Starting Scanner")
	s.Adapter.Connect()
	s.MotorTableDriver.Start()
	s.MotorOneCameraDriver.Start()
	s.MotorTwoCameraDriver.Start()
}
