package parser

import (
	"github.com/Ethereal-Coder/awesome-go-learn/spider/engine"
	"log"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	// compile := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
	compile := regexp.MustCompile(cityListRe)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(
			result.Items, string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				//ParserFunc: engine.NilParser,
				ParserFunc: ParseCity,
			})
		log.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	// log.Printf("Matches found: %d\n", len(matches))
	return result
}
