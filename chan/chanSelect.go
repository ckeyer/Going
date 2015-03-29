package main

import (
	"fmt"
	"time"
)

const (
	MAX_INDEX = 100000000
)

type CJStream struct {
	W      chan int
	Wnum   int
	Wcount chan int
	R      chan int
	Rnum   int
	Rcount chan int
}

func sleep(second int) {
	time.Sleep((time.Duration)(second) * time.Second)
}

func main() {

	fmt.Println("Start")
	start := time.Now().UnixNano()
	test1()
	end := time.Now().UnixNano()
	fmt.Println("Over")
	fmt.Println(end - start)
}

func test2() {
	count := 0
	for i := 1; i <= MAX_INDEX; i++ {
		count += i
	}
	fmt.Println("Count: ", count)
}
func test1() { // 16991094372
	cj := &CJStream{
		W:      make(chan int, 10),
		R:      make(chan int, 10),
		Wcount: make(chan int),
		Rcount: make(chan int),
		Wnum:   0,
		Rnum:   0,
	}
	go cj.Read()
	go cj.Write()
	for i := 1; i <= MAX_INDEX; i++ {
		select {
		case cj.W <- i:
		case cj.R <- i:
		}
	}
	cj.R <- 0
	cj.W <- 0
	// time.Sleep(time.Second)
	fmt.Println("RCount: ", cj.Rnum)
	fmt.Println("WCount: ", cj.Wnum)
	fmt.Println("Count: ", cj.Rnum+cj.Wnum)
}

func (this *CJStream) Read() {
	for v := range this.R {
		if v == 0 {
			this.Rcount <- this.Rnum
		} else {
			this.Rnum += v
			fmt.Println("R--", v)
		}
		if len(this.R) <= 0 { // 如果现有数据量为0，跳出循环
			continue
		}
	}
}

func (this *CJStream) Write() {
	for v := range this.W {
		if v == 0 {
			this.Wcount <- this.Wnum
		} else {
			this.Wnum += v
			fmt.Println("W--", v)
		}
		if len(this.W) <= 0 { // 如果现有数据量为0，跳出循环
			continue
		}
	}
}
