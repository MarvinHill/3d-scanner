package internal

import "sync"

type ScannerController struct {
	mu              sync.Mutex
	jobRunning      bool
	PhotosTaken     []Photo
	photosRequested []PhotoRequest
	scanner         *ScannerDriver
}

func NewJobScheduler(scanner *ScannerDriver) *ScannerController {
	return &ScannerController{
		jobRunning:      false,
		PhotosTaken:     make([]Photo, 0),
		photosRequested: make([]PhotoRequest, 0),
		scanner:         scanner,
	}
}

func (sc *ScannerController) StartScanJob(request []PhotoRequest) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	for _, photoRequest := range request {
		sc.PhotosTaken = append(sc.PhotosTaken, sc.scanner.TakePhoto(&photoRequest))
	}
}

func (sc *ScannerController) MoveByManualControl(control *ManualControlMessage) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if sc.jobRunning {
		return false
	}

	sc.scanner.MoveByManualControl(control)
	return true
}
