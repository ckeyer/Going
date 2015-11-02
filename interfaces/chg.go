package main

import (
	"fmt"
)

type AA struct {
	Id   int
	Name string
}

// NewAA ...
func NewAA(name string) interface{} {
	return &AA{
		Id:   123,
		Name: name,
	}
}

func main() {
	a := NewAA("ckkkk").(*AA)
	fmt.Printf("%#v \n", a)
}
