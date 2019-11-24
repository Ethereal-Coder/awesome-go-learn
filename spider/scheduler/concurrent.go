package scheduler

import (
	"github.com/Ethereal-Coder/awesome-go-learn/spider/engine"
)

type ConcurrentScheduler struct {
	workerChan chan engine.Request
}

func (s *ConcurrentScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *ConcurrentScheduler) Submit(r engine.Request) {
	// Send request down to worker chan
	go func() { s.workerChan <- r }()
}

func (s *ConcurrentScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
