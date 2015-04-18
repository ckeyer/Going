package main

import "fmt"
import (
	"runtime"
	"time"
)

func foo(c chan<- int) {
	fmt.Println("World")
	time.Sleep(time.Second * 2)
	fmt.Println("Hello")
	c <- 2
}

func sum(x, y int, c chan int) {
	// time.Sleep(time.Second * 1)
	fmt.Println(<-c)
	// c <- x + y
}

func main() {
	c := make(chan int)
	go foo(c)
	go sum(24, 18, c)
	// c <- 3
	runtime.Gosched()
	fmt.Println("over")
	// time.Sleep(time.Second * 3)
}

func write(c chan<- []byte) {
	time.Sleep(time.Millisecond * 500)
	var bs []byte
	copy(bs, "abcde")
	c <- bs
}
func read(c <-chan []byte) {

}
