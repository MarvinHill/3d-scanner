package internal

import (
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

type ScannerDriver struct {
	Adapter          *raspi.Adaptor
	MotorOneDriver   *gpio.StepperDriver
	MotorTwoDriver   *gpio.StepperDriver
	MotorThreeDriver *gpio.StepperDriver
}

func (s *ScannerDriver) Run() {
	s.Adapter = raspi.NewAdaptor()
	s.MotorOneDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"8", "9", "7", "0"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorTwoDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"2", "3", "12", "13"}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorThreeDriver = gpio.NewStepperDriver(s.Adapter, [4]string{"15", "16", "1", "4"}, gpio.StepperModes.SinglePhaseStepping, 2048)
}
