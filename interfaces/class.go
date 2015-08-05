package main

import (
	"fmt"
)

type A struct {
	Name string
}

func (a *A) String() string {
	return "AA:" + a.Name
}
func (a *A) Set() {
	a.Name = "AAAAAAAAAAA"
}

type B struct {
	A
}

func (b *B) Set() {
	b.Name = "BBBBBBBBBBB"
}

type C struct {
	B
}

func (c *C) String() string {
	return "CC:" + c.Name
}

func main() {
	a := &A{}
	a.Set()
	fmt.Println(a)
	b := &B{}
	b.Set()
	fmt.Println(b)
	c := &C{}
	c.Set()
	fmt.Println(c)
}
