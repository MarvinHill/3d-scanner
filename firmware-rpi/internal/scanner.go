package internal

import (
	"fmt"
	"sync"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

type ScannerDriver struct {
	mu                   sync.Mutex
	level                bool
	Adapter              *raspi.Adaptor
	MotorTableDriver     *gpio.StepperDriver
	MotorOneCameraDriver *gpio.StepperDriver
	MotorTwoCameraDriver *gpio.StepperDriver
	manualStepAmount     int
	ManualControl        *ManualControl
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
	s.ManualControl = nil
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

func (s *ScannerDriver) SetManualControl(control *ManualControl) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ManualControl = control
}

func (s *ScannerDriver) Run() {
	fmt.Println("Starting Scanner")
	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			fmt.Println("Move")

			if s.ManualControl != nil {
				switch s.ManualControl.MoveType {
				case "tableAxisPlus":
					s.MotorTableDriver.Move(s.manualStepAmount)
					return
				case "tableAxisMinus":
					s.MotorTableDriver.Move(-s.manualStepAmount)
					return
				case "cameraAxisPlus":
					s.MotorOneCameraDriver.Move(s.manualStepAmount)
					s.MotorTwoCameraDriver.Move(-s.manualStepAmount)
					return
				case "cameraAxisMinus":
					s.MotorOneCameraDriver.Move(-s.manualStepAmount)
					s.MotorTwoCameraDriver.Move(s.manualStepAmount)
					return
				}
			}

		})
	}

	robot := gobot.NewRobot("Scanner roboter",
		[]gobot.Connection{s.Adapter},
		[]gobot.Device{s.MotorTableDriver, s.MotorOneCameraDriver, s.MotorTwoCameraDriver},
		work,
	)

	robot.Start()
}
