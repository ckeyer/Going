package turl

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParseUrl(t *testing.T) {
	u, err := url.Parse("192.168.1.2:233")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", u.String())

	uu := &url.URL{}
	uu.Query().Set("a", "aaa")
	uu.Query().Set("b", "bbb")
	fmt.Println(uu.Query().Encode())
	fmt.Println(uu.String())

}
