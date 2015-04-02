package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start...")
	// ai := time.Now().UnixNano()
	t := time.Unix(time.Now().Unix(), 0)
	fmt.Println(t.String())
	fmt.Println(time.Now().Unix())
	fmt.Println("Over")
}
