package scheduler

import "learnGo/crawler/engine"

type SimplerScheduler struct {
	workerChan chan engine.Request
}

func (s *SimplerScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimplerScheduler) Submit(r engine.Request) {
	// send requests down to worker chan
	go func() {
		s.workerChan <- r
	}()
}
