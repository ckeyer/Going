package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//  返回Get数据
func HttpGet(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}

var url string = `http://zhangmenshiting.baidu.com/data2/music/134379002/1215232761426935661.mp3?xcode=1905f93142116011914a4383e913a546c650e349c3915883`

func main() {
	c, _ := HttpGet(url)
	fmt.Println(c)
}
