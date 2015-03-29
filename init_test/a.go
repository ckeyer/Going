package main

import (
	"fmt"
)

type People struct {
	Id   int64  `json:"ID"`
	Name string `json:"NAME"`
}

func main() {
	fmt.Println("Start")

	var b uint8 = 0xa
	fmt.Printf("%b \n", 254)

	fmt.Println("Over")
}
