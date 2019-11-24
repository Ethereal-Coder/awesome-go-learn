package engine

import (
	"log"
)

type QueuedEngine struct {
	Scheduler   IQueued
	WorkerCount int
}

func (e *QueuedEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func (e *QueuedEngine) createWorker(in chan Request, out chan ParseResult, s IQueued) {
	//in := make(chan Request)
	go func() {
		for {
			// Tell scheduler i'm ready
			s.WorkerReady(in)

			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
