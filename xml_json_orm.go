package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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

}
