package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "unicode/utf8"
)

var (
	url = "http://api.map.baidu.com/location/ip?ak=F454f8a5efe5e577997931cc01de3974&ip="
	ips = []string{"182.100.67.112",
		"186.121.210.50",
		"187.141.5.177",
		"187.178.206.67",
		"188.138.72.27",
		"182.100.67.113",
		"182.100.67.114",
		"183.61.109.240",
		"183.101.81.148",
		"184.168.119.160",
		"189.112.5.152"}
)

type IP2Address struct {
	Address string
	Content AddressContent
}
type AddressContent struct {
	Address_detail AddressDetail
}
type AddressDetail struct {
	Province  string
	City      string
	District  string
	City_code int
}

func main() {
	for i, v := range ips {
		fmt.Printf("%d: %s\t %s\n", i, v, GetAddress(v))
	}
}

func GetAddress(ip string) (ip2addr IP2Address) {
	data, s := HttpGet(url + ip)
	if s == 200 {

		e := json.Unmarshal(data, &ip2addr)
		if e != nil {
			fmt.Println(e)
			return
		}
	}
	return
}

//  返回Get数据
func HttpGet(url string) (data []byte, statusCode int) {
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
	return
}
