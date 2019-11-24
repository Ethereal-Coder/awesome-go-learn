package engine

type scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type IConcurrent interface {
	scheduler
}

type IQueued interface {
	scheduler
	WorkerReady(chan Request)
}
