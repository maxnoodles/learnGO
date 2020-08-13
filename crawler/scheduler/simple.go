package scheduler

import "learnGo/crawler/engine"

type SimplerScheduler struct {
	workerChan chan engine.Request
}

func (s *SimplerScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimplerScheduler) WorkerReady(requests chan engine.Request) {
}

func (s *SimplerScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimplerScheduler) Submit(r engine.Request) {
	// send requests down to worker chan
	go func() {
		s.workerChan <- r
	}()
}
