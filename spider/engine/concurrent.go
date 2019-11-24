package engine

type ConcurrentEngine struct {
	Scheduler   IConcurrent
	WorkerCount int
	ItemChan    chan Item
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	e.Scheduler.Run()
	out := make(chan ParseResult)

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("Got item #%d: %v", itemCount, item)
			//itemCount++
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
