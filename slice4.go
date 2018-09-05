package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {
	var s1 = make([]int, 3, 15)
	s2 := make([]int, 6, 60)

	printSlice(s1)
	printSlice(s2)
}
