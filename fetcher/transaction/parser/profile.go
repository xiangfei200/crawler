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

func parseProfile(contents []byte,name string) engine.ParseResult {
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
	result := engine.ParseResult{
		Items:    []interface{}{profile},
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
