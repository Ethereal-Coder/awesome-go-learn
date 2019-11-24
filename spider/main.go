package main

import (
	"github.com/Ethereal-Coder/awesome-go-learn/spider/engine"
	"github.com/Ethereal-Coder/awesome-go-learn/spider/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
