package engine

type ParserFunc func(
	contents []byte, url string) ParseResult

type Request struct {
	Url string
	// 关联挂在下一步操作，并拿到相关数据，进行赋值
	ParserFunc
}


type ParseResult struct {
	Requests []Request
	Items []Item
}

//定义通用字段并拓展字段 使用interface可以存放任意类型字段
//为了配合elasticsearch 我们想把Type和Id自定义，Index由配置文件配置
type Item struct {
	Url string
	Type string
	Id string
	Payload interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}