package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func sleep(second int) {
	time.Sleep((time.Duration)(second) * time.Second)
}

func main() {
	test3()
	// test2()
	// test1()
}

func test3() {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case c <- 1:
		case c <- 2:
		case c <- 3:
		case c <- 4:
		}
	}
}

func test2() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("hello: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
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
