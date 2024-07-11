package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

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

const (
	motorTable1 = "3"
	motorTable2 = "5"
	motorTable3 = "7"
	motorTable4 = "11"

	motorOne1 = "23"
	motorOne2 = "29"
	motorOne3 = "31"
	motorOne4 = "33"

	motorTwo1 = "8"
	motorTwo2 = "10"
	motorTwo3 = "12"
	motorTwo4 = "16"

	MaxMotorSteps = 2048.0
	MaxDeg        = 360.0
)

// Steps to Degres
// 2048 steps = 360 degrees
func StepsToDegrees(steps int) float64 {
	return float64(steps) * (MaxDeg / MaxMotorSteps)
}

func NewScannerDriver() *ScannerDriver {
	s := &ScannerDriver{}
	s.Adapter = raspi.NewAdaptor()
	s.MotorTableDriver = gpio.NewStepperDriver(s.Adapter, [4]string{motorTable1, motorTable2, motorTable3, motorTable4}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorOneCameraDriver = gpio.NewStepperDriver(s.Adapter, [4]string{motorOne1, motorOne2, motorOne3, motorOne4}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.MotorTwoCameraDriver = gpio.NewStepperDriver(s.Adapter, [4]string{motorTwo1, motorTwo2, motorTwo3, motorTwo4}, gpio.StepperModes.SinglePhaseStepping, 2048)
	s.manualStepAmount = 1
	return s
}

func (s *ScannerDriver) LevelAll() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.levelAll()
}

func (s *ScannerDriver) LevelSites() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.levelSites()
}

func (s *ScannerDriver) levelAll() {
	fmt.Println("Leveling Scanner")
	s.level()
	s.CurrentPosition = NewPosition(0, 0)
}

func (s *ScannerDriver) levelSites() {
	fmt.Println("Leveling Scanner CameraAxis")
	s.level()
	s.CurrentPosition = NewPosition(0, s.CurrentPosition.TableAxis)
}

func (s *ScannerDriver) level() {
	fmt.Println("level")
	oneLevel := false
	twoLevel := false

	for oneLevel == false || twoLevel == false {
		fmt.Println("keep leveling")
		time.Sleep(50 * time.Millisecond)
		// check level of axis motors

		if oneLevel == false {
			oneFirstCheck, _ := s.Adapter.DigitalRead("36")
			if oneFirstCheck == 1 {
				fmt.Println("Scanner Motor one leveled")
				oneLevel = true
			} else {
				// move down
				fmt.Println("Moving camera axis 1")
				s.MotorOneCameraDriver.Move(2)
				oneSecondCheck, _ := s.Adapter.DigitalRead("36")
				if oneSecondCheck == 1 {
					oneLevel = true
				}
			}

		}

		if twoLevel == false {
			twoFirstCheck, _ := s.Adapter.DigitalRead("38")
			if twoFirstCheck == 1 {
				fmt.Println("Scanner Motor one leveled")
				twoLevel = true
			} else {
				// move down
				fmt.Println("Moving camera axis 2")
				s.MotorTwoCameraDriver.Move(2)
				twoSecondCheck, _ := s.Adapter.DigitalRead("38")
				if twoSecondCheck == 1 {
					twoLevel = true
				}
			}
		}

	}
	fmt.Println("Scanner leveled")
	s.ResetMotors()
}

func (s *ScannerDriver) TakePhoto(request PhotoRequest) (Photo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// if s.CurrentPosition.TableAxis != request.ToPosition().TableAxis {
	// 	s.levelSites()
	// }

	requestedPos := request.ToPosition()
	tableAxisDiff := requestedPos.TableAxis - s.CurrentPosition.TableAxis // invert because of stepper motor mounting direction
	cameraAxisDiff := requestedPos.CameraAxis - s.CurrentPosition.CameraAxis

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Moving camera axis 1")
	go func() {
		s.MotorOneCameraDriver.MoveDeg(-cameraAxisDiff * 2)
		wg.Done()
	}()
	fmt.Println("Moving camera axis 2")
	go func() {
		go s.MotorTwoCameraDriver.MoveDeg(-cameraAxisDiff * 2)
		wg.Done()
	}()
	fmt.Println("Moving table")
	s.MotorTableDriver.MoveDeg(tableAxisDiff)
	wg.Wait()
	s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, NewPosition(cameraAxisDiff, tableAxisDiff))

	// s.takePhoto()

	cmd := exec.Command("rpicam-still", "-o", "/var/images/image.jpg")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error taking photo")
		return Photo{}, errors.New("Error taking photo")
	}

	fileData, err := os.ReadFile("/var/images/image.jpg")

	if err != nil {
		fmt.Println("Error loading photo as byte array")
		return Photo{}, errors.New("Error taking photo")
	}

	fmt.Println("Photo taken")

	return Photo{
		AngleCameraAxis: s.CurrentPosition.CameraAxis,
		AngleTableAxis:  s.CurrentPosition.TableAxis,
		PhotoData:       string(fileData),
	}, nil
}

func (s *ScannerDriver) MoveByManualControl(movement string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println("Moving by manual control")

	switch movement {
	case "c_pl":
		prevPos := s.CurrentPosition
		s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, Position{CameraAxis: s.manualStepAmount, TableAxis: 0})
		s.CurrentPosition.Print()
		if prevPos.Equals(s.CurrentPosition) {
			break
		}
		fmt.Println("Moving camera axis 1")
		go func() {
			s.MotorOneCameraDriver.MoveDeg(-s.manualStepAmount * 2)
		}()
		fmt.Println("Moving camera axis 2")
		go func() {
			s.MotorTwoCameraDriver.MoveDeg(-s.manualStepAmount * 2)
		}()
		break
	case "c_min":
		prevPos := s.CurrentPosition
		s.CurrentPosition = AddMovementToPosition(s.CurrentPosition, Position{CameraAxis: -s.manualStepAmount, TableAxis: 0})
		s.CurrentPosition.Print()
		if prevPos.Equals(s.CurrentPosition) {
			break
		}
		fmt.Println("Moving camera axis 1")
		go func() {
			go s.MotorOneCameraDriver.MoveDeg(s.manualStepAmount * 2)
		}()
		fmt.Println("Moving camera axis 2")
		go func() {
			go s.MotorTwoCameraDriver.MoveDeg(s.manualStepAmount * 2)
		}()
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

	s.MotorTableDriver.SetSpeed(20)
	s.MotorOneCameraDriver.SetSpeed(10)
	s.MotorTwoCameraDriver.SetSpeed(10)

	s.ResetMotors()

	// Sensor 1
	s.Adapter.DigitalWrite("35", 1)
	// Sensor 2
	s.Adapter.DigitalWrite("37", 1)
}

func (s *ScannerDriver) ResetMotors() {
	fmt.Println("Stopping Scanner")
	s.MotorTableDriver.Sleep()
	s.MotorOneCameraDriver.Sleep()
	s.MotorTwoCameraDriver.Sleep()
}

func (s *ScannerDriver) SetScannerLevel() {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println("Setting Scanner Level")
	s.CurrentPosition = NewPosition(0, 0)
}
