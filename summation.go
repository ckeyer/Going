package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	MAX_NUM  = 100 //0 * 10000  * 10000
	MAX_CPUS = 10
)

func main() {
	runtime.GOMAXPROCS(MAX_CPUS)
	var mycount [MAX_CPUS]MyCount
	var i int64
	for i = 0; i < MAX_CPUS; i++ {
		mycount[i] = NewMyCount(i)
		defer close(mycount[i].CCount)
		go mycount[i].Add(index)
	}
	// start := time.Now().UnixNano()
	// fmt.Println(MAX_NUM)
	// end := time.Now().UnixNano()
	// fmt.Println(end - start)
	for i := 0; i < MAX_NUM; i++ {
		select {
		case c[0] <- i:
		case c[1] <- i:
		case c[2] <- i:
		case c[3] <- i:
		case c[4] <- i:
		case c[5] <- i:
		case c[6] <- i:
		case c[7] <- i:
		case c[8] <- i:
		case c[9] <- i:
		}
	}
	time.Sleep(2 * time.Second)
}

type MyCount struct {
	Index  int64
	Count  int64
	CCount chan int64
}

func NewMyCount(i int64) (m *MyCount) {
	m = &MyCount{
		Index:  i,
		Count:  0,
		CCount: make(chan int64),
	}
	return
}
func (this *MyCount) Add(index chan int) {

}
