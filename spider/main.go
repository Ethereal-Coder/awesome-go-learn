package main

import (
	"github.com/Ethereal-Coder/awesome-go-learn/spider/engine"
	"github.com/Ethereal-Coder/awesome-go-learn/spider/persist"
	"github.com/Ethereal-Coder/awesome-go-learn/spider/scheduler"
	"github.com/Ethereal-Coder/awesome-go-learn/spider/zhenai/parser"
)

func main() {
	// single task
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// concurrent schedule
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.ConcurrentScheduler{},
	//	WorkerCount: 100,
	//}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// queued scheduler
	//e := engine.QueuedEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 100,
	//}
	//e.Run(engine.Request{
	//	//Url:        "http://www.zhenai.com/zhenghun",
	//	//ParserFunc: parser.ParseCityList,
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})

	// Page
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.ConcurrentScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
		//Url:        "http://www.zhenai.com/zhenghun",
		//ParserFunc: parser.ParseCityList,
	})
}
