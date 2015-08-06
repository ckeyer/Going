package main

import "fmt"
import "net/http"
import "github.com/PuerkitoBio/goquery"

var (
	url  = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxcheckurl?requrl=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMjM5NDE0MjI4MA%3D%3D%26mid%3D208304723%26idx%3D4%26sn%3Db429429589e336e8eac4c98118f3a4c4%26scene%3D2%26from%3Dtimeline%26isappinstalled%3D0%23rd&skey=%40crypt_13e32360_eeb8fd095ab87dd57fe31605a7daa4f6&deviceid=e582064974121749&pass_ticket=KgjqHMPrkf4O7sVnhzVXfbiZUC4tt%252FVwTOF7jzR6wYk8ZxBA0nZjhanT6r7y%252BoJG&opcode=2&scene=1&username=@e8ebbd1b58416ce58442cf0bfb6e94ff3306a22d27020d02e6c7abe07fa9c11f"
	url2 = "http://mp.weixin.qq.com/s?__biz=MjM5NDE0MjI4MA==&mid=208304723&idx=4&sn=b429429589e336e8eac4c98118f3a4c4&scene=2&from=timeline&isappinstalled=0&key=0acd51d81cb052bc18293c56365f5239cdb341cc82a47c6824ee73384b195af410d01fa3c93c0dad10f5afd7eacaa092&ascene=1&uin=MTY1NzU5MDg4Mw%3D%3D&devicetype=webwx&version=70000001&pass_ticket=KgjqHMPrkf4O7sVnhzVXfbiZUC4tt%2FVwTOF7jzR6wYk8ZxBA0nZjhanT6r7y%2BoJG"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

}

// test get
func test() (err error) {
	// cli:=&http.Client{	}
	res, err := http.Get(url2)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	fmt.Println(doc.Text())
	doc.Find(".rich_media_content").Each(func(i int, s *goquery.Selection) {
		s.Find("p").Each(func(i int, ss *goquery.Selection) {
			fmt.Println(ss.Text())
		})
	})

	return
}
