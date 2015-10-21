package main

import "fmt"

func slice() []int {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{-1, -2, -3}
	return append(append(s1[:1], s2...), s1[1:]...)
}

func slice1() []int {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{-1, -2, -3, -4}
	return append(append(s1[:1], s2...), s1[1:]...)
}

func main() {

	// 输出结果？
	fmt.Printf("func slice=%+v\n", slice())
	fmt.Printf("func slice1=%+v\n", slice1())
}
