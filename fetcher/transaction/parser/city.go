package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, value := range all {
		// 拷贝用户名  否则 ParserFunc 始终会是最后一个值，
		//因为外部变量闭包中相当于全局变量，变量不会被释放或删除，因此到执行(被调用)时是最后一个值。
		//拷贝的话，因为每次都是重新赋值:=，属于局部变量。
		//https://zhuanlan.zhihu.com/p/92634505 闭包的延迟绑定
		name := string(value[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(value[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return parseProfile(contents,name)
			},
		})
		limit--
		if limit == 0{
			break
		}
	}
	//fmt.Printf("Matched found;%d\n",len(all))
	return result
}
