package parser

import (
	"github.com/Ethereal-Coder/awesome-go-learn/spider/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)
)

func ParseCity(contents []byte) engine.ParseResult {
	rs := engine.ParseResult{}
	matches := profileRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		name := string(m[2])
		rs.Items = append(rs.Items, "User "+string(m[2]))

		rs.Requests = append(rs.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})

	}

	return rs
}
