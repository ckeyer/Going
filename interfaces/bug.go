package main

import (
	"fmt"
)

type CkString string

func (c *CkString) String() string {
	return fmt.Sprint(c)
}

func main() {
	var c CkString = "haha"
	c.String()
}
