package internal

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

type PhotoRequest struct {
	angleCameraAxis int
	angleTableAxis  int
}

type PhotoJob struct {
	PhotoRequest *[]PhotoRequest
}

type ScannerDriver struct {
	Adapter          *raspi.Adaptor
	MotorOneDriver   *gpio.StepperDriver
	MotorTwoDriver   *gpio.StepperDriver
	MotorThreeDriver *gpio.StepperDriver
	currentPosition  []int
	PhotoJobs        chan PhotoJob
	JobResultChannel chan []byte
}

func NewScannerDriver() *ScannerDriver {
	s := &ScannerDriver{}
	s.Adapter = raspi.NewAdaptor()
	s.MotorOneDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"3", "5", "7", "11"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorTwoDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"13", "15", "19", "21"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorThreeDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"8", "10", "12", "16"}, gpio.StepperModes.SinglePhaseStepping, 2048)

	s.PhotoJobs = make(chan PhotoJob, 100)
	s.JobResultChannel = make(chan []byte)

	return s
}

func (s *ScannerDriver) LevelSites() {
	// Level the scanner Sites
}

func (s *ScannerDriver) Run() {
	fmt.Println("Starting Scanner")
	work := func() {

		gobot.Every(1*time.Second, func() {
			fmt.Println("Move")
			s.MotorOneDriver.MoveDeg(30)
			s.MotorTwoDriver.MoveDeg(30)
			s.MotorThreeDriver.MoveDeg(30)
		})
	}

	robot := gobot.NewRobot("Scanner roboter",
		[]gobot.Connection{s.Adapter},
		[]gobot.Device{s.MotorOneDriver, s.MotorTwoDriver, s.MotorThreeDriver},
		work,
	)

	robot.Start()

}
