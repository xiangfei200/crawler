package parser

import (
	"crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
// 只解析城市，后续的由城市自行调闭包，再调出详细的人
var realteRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte,_ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, value := range matches {
		// 拷贝用户名  否则 ParserFunc 始终会是最后一个值，
		//因为外部变量闭包中相当于全局变量，变量不会被释放或删除，因此到执行(被调用)时是最后一个值。
		//拷贝的话，因为每次都是重新赋值:=，属于局部变量。
		//https://zhuanlan.zhihu.com/p/92634505 闭包的延迟绑定
		name := string(value[2])
		url := string(value[1])
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:       url ,
			//ParserFunc: func(contents []byte) engine.ParseResult {
			//	return parseProfile(contents,url,name)
			//},
			//函数调用，是引用值，会自动copy。上面的name := string(value[2])都可以去掉，直接用string(value[2])
			Parser: NewProfileParser(name),
		})
	}
	matches = realteRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: engine.NewFuncParser(ParseCity,"ParseCity"),
		})
	}
	//fmt.Printf("Matched found;%d\n",len(all))
	return result
}

