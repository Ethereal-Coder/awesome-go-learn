package scheduler

import (
	"github.com/Ethereal-Coder/awesome-go-learn/spider/engine"
)

type ConcurrentScheduler struct {
	workerChan chan engine.Request
}

func (s *ConcurrentScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *ConcurrentScheduler) Submit(r engine.Request) {
	// Send request down to worker chan
	go func() { s.workerChan <- r }()
}
