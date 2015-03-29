package main

import (
	"fmt"
	"time"
)

func sleep(second int) {
	time.Sleep((time.Duration)(second) * time.Second)
}

func main() {
	test1()
}

func test1() {
	c := make(chan int, 1) //
	// defer
	go test1ChanGet(c)
	go test1ChanSet(c, 11)
	go test1ChanSet(c, 22)
	go test1ChanSet(c, 44)
	sleep(3)
	go test1ChanGet(c)
	go test1ChanSet(c, 33)
	// fmt.Println(c)
	sleep(5)
	close(c)
	if c != nil {
		fmt.Println("sdf")
		for e := range c {
			fmt.Println("end : ", e)
		}
	}

}

func test1ChanGet(c <-chan int) {

	sleep(2)
	fmt.Println("get :", <-c)
}
func test1ChanSet(c chan<- int, ci int) {
	fmt.Println("before set:", ci)
	c <- ci

	fmt.Println("set over:", ci)
}
