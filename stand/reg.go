package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Hello")
	f, err := os.OpenFile("map.md", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Over")
		return
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("OverFile")
	}
	regs := regexp.MustCompile(`<!--`)
	rege := regexp.MustCompile(`-->`)
	start := regs.FindStringIndex(string(bs))
	end := rege.FindStringIndex(string(bs))
	if len(start) < 2 || len(end) < 2 {
		return
	}
	fmt.Println(start[1], end[0])
	fmt.Printf("%s\n", bs[start[1]:end[0]])
	fmt.Printf("%s\n", bs[end[1]:])
}
