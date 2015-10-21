package main

import "fmt"

type T struct {
	A int
	B int
}

func main() {
	fmt.Println("Hello")
	m := make(map[int]string, 3)
	fmt.Printf("%#v\n", m)
	fmt.Println(len(m))
}
