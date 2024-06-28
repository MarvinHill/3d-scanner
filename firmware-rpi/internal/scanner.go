package internal

import (
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
	s.MotorOneDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"8", "9", "7", "0"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorTwoDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"2", "3", "12", "13"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorThreeDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"15", "16", "1", "4"}, gpio.StepperModes.SinglePhaseStepping, 2048)

	s.PhotoJobs = make(chan PhotoJob, 100)
	s.JobResultChannel = make(chan []byte)

	return s
}

func (s *ScannerDriver) LevelSites() {
	// Level the scanner Sites
}

func (s *ScannerDriver) Run() {
	for true {
		//Check if there are jobs in the queue
	}

}
