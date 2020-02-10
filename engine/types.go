package engine

type Request struct {
	Url string
	// 关联挂在下一步操作，并拿到相关数据，进行赋值
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}