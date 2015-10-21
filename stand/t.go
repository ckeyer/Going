package main

import "fmt"

type T struct {
	Index int64
	Value string
}

func main() {
	t1 := new(T)
	t2 := &T{}
	t3 := &T{0, ""}
	t4 := &T{Index: 0, Value: ""}

	fmt.Printf("%#v\n", t1)
	fmt.Printf("%#v\n", t2)
	fmt.Printf("%#v\n", t3)
	fmt.Printf("%#v\n", t4)
}
