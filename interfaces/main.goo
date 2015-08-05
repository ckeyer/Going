package main

import (
	"fmt"
)

// type RouterString func() string
// type RouterBytes func() []byte

// func AddRouter(router interface{}) {
// 	router := (*router).(type
// 	switch router {
// 	case RouterString:
// 		println("string")
// 	case RouterBytes:
// 		println("bytes")
// 	default:
// 		println("unknown types")
// 	}
// }
type User struct {
	Id   int
	Name string
}
type Admin struct {
	Id      int
	Name    string
	Company string
}
type Custom struct {
	Id    int
	Name  string
	sales float32
}

func doTest1(a interface{}) {
	// t := a.(type)
	switch t := a.(type) {
	case User:
		fmt.Println("User")
		tt, ok := a.(User)
		if ok {
			fmt.Println(t, tt, ok)
		}
	case Admin:
		fmt.Println("Admin")
		tt, ok := a.(Admin)
		if ok {
			fmt.Println(t, tt, ok)
		}
	case Custom:
		fmt.Println("Custom")
		tt, ok := a.(Custom)
		if ok {
			fmt.Println(t, tt, ok)
		}
	default:
		fmt.Println(t)
	}
}
func test1() {
	a := &Admin{Id: 10, Name: "Wang"}
	doTest1(*a)
}
func main() {
	test1()
}
