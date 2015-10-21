package main

import "fmt"
import "strings"
import "runtime"

func main() {
	strs := []string{}
	for i := 0; i < 10; i++ {
		strs = append(strs, fmt.Sprint(i, "__"))
	}
	fmt.Println("`" + strings.Join(strs, "`,`") + "`")
	fmt.Println(runtime.CPUProfile())
}
