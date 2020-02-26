package engine

type ParserFunc func(
	contents []byte, url string) ParseResult

//定义网络传输结构体必须的函数，继承这个接口的所有结构体必须实现这些函数
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string,args interface{})
}

type Request struct {
	Url string
	// 关联挂在下一步操作，并拿到相关数据，进行赋值
	//原来是个ParserFunc 在网络无法传输，不能解析成json报文，因此需要我们把解析器自行序列化、反序列化
	//
	Parser Parser
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

type NilParse struct {
}

func (NilParse) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParse) Serialize() (name string, args interface{}) {
	return "NilParser",nil
}

//func NilParse([]byte) ParseResult {
//	return ParseResult{}
//}

//定义网络传输所需要的结构体，以方便转成json报文
type FuncParser struct {
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents,url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name,nil
}

func NewFuncParser(p ParserFunc,name string) *FuncParser{
	return &FuncParser{
		parser: p,
		name:   name,
	}
}