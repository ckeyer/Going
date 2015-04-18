package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	USE_GOROUTINES = 5
	MAX_INDEX      = 100000000000
)

func main() {
	runtime.GOMAXPROCS(USE_GOROUTINES)
	s := time.Now().UnixNano()
	var Count_all int64 = 0
	ch := make(chan int64, 10)
	var i int64 = 1
	for ; i <= USE_GOROUTINES; i++ {
		go GoRun(i, USE_GOROUTINES, ch)
	}
	return_times := 0
	for {
		if return_times >= USE_GOROUTINES {
			break
		}
		select {
		case count_one := <-ch:
			return_times++
			Count_all += count_one
		default:
		}
	}
	e := time.Now().UnixNano()
	fmt.Println("ALL_COUNT :", Count_all)
	fmt.Println((e - s))
	time.Sleep(time.Second * 2)
	s2 := time.Now().UnixNano()
	var Count_all2 int64 = 0
	for i = 1; i < MAX_INDEX; i++ {
		Count_all2 += i
	}
	e2 := time.Now().UnixNano()
	fmt.Println("ALL_COUNT :", Count_all2)
	fmt.Println((e2 - s2))
}

func GoRun(first, step int64, ch chan int64) {
	var count int64 = 0
	for i := first; i <= MAX_INDEX; i += step {
		count += i
	}
	ch <- count
}

type CKCount struct {
	count int64
	Node  chan int64
}

func (this *CKCount) Add() {

}

func (this *CKCount) GetCount() int64 {
	return this.count
}
