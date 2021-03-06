package engine

import (
	"github.com/Ethereal-Coder/awesome-go-learn/spider/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (engine SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, e := worker(r)
		if e != nil {
			continue
		}

		//requests = append(requests, parseResult.Requests...)

		for _, r := range parseResult.Requests {
			if isDuplicate(r.Url) {
				continue
			}
			requests = append(requests, r)
		}

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: err fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
