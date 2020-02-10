package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRexp = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^C<]+)</a>`

/**
解析获取到的内容，转化成下一步的request和此页的data
*/
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRexp)

	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 5
	for _, value := range all {
		// 城市名
		result.Items = append(result.Items, "City "+string(value[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(value[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0{
			break
		}
	}
	//fmt.Printf("Matched found;%d\n",len(all))
	return result
}
