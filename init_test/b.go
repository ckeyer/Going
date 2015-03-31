package main

import (
	"fmt"
)

func main() {
	bs := make([]byte, 8)

	copy(bs, []byte("1234"))

	copy(bs[4:], []byte("5678"))

	fmt.Printf("%s\n", bs)

}

// import logpkg "github.com/funxdata/log4go"

// var (
// 	log logpkg.Logger
// )

// func init() {
// 	log = logpkg.NewConsoleLogger(logpkg.WARNING)
// }
// func main() {
// 	log.Error("hh")
// 	for i := 1; i <= 100; i++ {
// 		log.Info(i, " ")
// 		if 0 == i%3 {
// 			log.Fine("Fizz")
// 		}
// 		if 0 == i%5 {
// 			log.Warn("Buzz")
// 		}
// 		log.Warn("")
// 	}
// }
