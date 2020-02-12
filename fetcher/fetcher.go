package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//定义请求间隔，防止服务器方的禁止请求频率过快操作
//100毫秒
var rateLimiter = time.Tick(100*time.Millisecond)

/**
输入一个url 输出byte
 */
func Fetch(url string) ([]byte,error){
	<-rateLimiter
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	//resp, err := http.Get(url)
	//if err != nil{
	//	return nil,err
	//}
	defer resp.Body.Close()
	//fmt.Printf("%s",resp.StatusCode)
	if resp.StatusCode != http.StatusOK{
		//fmt.Println("Error: status code-",resp.StatusCode)
		return nil,fmt.Errorf("wrong status code:%d",resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// Peek 是针对bufio.NewReader(r)的 先读了1024个字节，读完之后缓存起来，后面再用NewReader读的话就从1025开始读了
	//bytes, err := bufio.NewReader(r).Peek(1024)
	bytes, err := r.Peek(1024)
	if err !=nil{
		//读取不出来文件的1024字节的话，默认是utf8
		return unicode.UTF8
	}
	//根据前1024个bytes来判断编码
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}


