package main

import (
	"fmt"
	"net/url"
)

func main() {

	// var srcstr = make([]byte, 20)
	// fmt.Println(base64.URLEncoding.EncodeToString([]byte("【百万吨】注册【s】")))

	fmt.Println(url.QueryEscape("URL编码解码"))
	// fmt.Println(srcstr)
}
