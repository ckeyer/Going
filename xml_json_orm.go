package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type TestUser struct {
	Id    int64      `json:"-",xml:"-"`
	Name  string     `json:"name",xml:"Name",orm:"size(16)"`
	Class *TestClass `orm:"one"`
}
type TestClass struct {
	Id   int64  `json:"-",xml:"-"`
	Name string `json:"name",xml:"Name",orm:"size(16)"`
}

func main() {
	class1 := &TestClass{
		Name: "Math",
	}

	user1 := &TestUser{
		Name:  "ckeyer",
		Class: class1,
	}
	bsjson, _ := json.Marshal(user1)
	bsxml, _ := xml.Marshal(user1)

	fmt.Printf("%s\n", bsjson)
	fmt.Printf("%s\n", bsxml)

}
