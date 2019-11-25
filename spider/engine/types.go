package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string      `db:"id"`
	Url     string      `db:"url"`
	Type    string      `db:"type"`
	Payload interface{} `db:"payload"`
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
