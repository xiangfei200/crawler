package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

func Transfer(newstr string) string {
	return strconv.QuoteToASCII(newstr)

}

//var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-[a-z0-9]>([\s\d!&@#$])</h1>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]{2,3})岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]{2,3})cm</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>`+"(未婚|离异|丧偶)"+`</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:`+"(\\d+(\\.\\d{1,2})?-\\d+(\\.\\d{1,2})?[百|千|万|百万])"+`</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>[\p{Han}]*`+"(小学|初中|中专|高中|高专|大专|本科|硕士|博士|博士后)"+`</div>`)
var hukouRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>`+"籍贯:([^<]*)"+`</div>`)
var idUrlRe = regexp.MustCompile(`http[s]?://[a-z]+.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte,url string,name string) engine.ParseResult {
	profile := model.Profile{}
	age, err := strconv.Atoi(extraString(contents, ageRe))
	if err == nil{
		profile.Age = age
	}
	height, err := strconv.Atoi(extraString(contents, heightRe))
	if err == nil{
		profile.Height = height
	}
	profile.Name = name
	profile.Marriage = extraString(contents,marriageRe)
	profile.Income = extraString(contents,incomeRe)
	profile.Education = extraString(contents,educationRe)
	profile.Hukou = extraString(contents,hukouRe)
	//fmt.Printf("%v",profile)
	var result = engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extraString([]byte(url),idUrlRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extraString(contents []byte,re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
		return string(match[1])
	}else{
		return ""
	}
}

//专门为从用户列表获取用户信息而创建的结构体，这个parse多了个参数，需要中转一下
type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents,url,p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser",p.userName
}

//func NewProfileParser(name string) engine.ParserFunc{
//	return func(content []byte, url string) engine.ParseResult {
//		return parseProfile(content,url,name)
//	}
//}

//专门为从用户列表获取用户信息而创建的工厂、构造函数
func NewProfileParser(name string) *ProfileParser{
	return &ProfileParser{userName:name}
}