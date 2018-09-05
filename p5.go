package main

import (
	"fmt"
)

func modify1(arr *[3]int) {
	(*arr)[0] = 10
}

func modify2(sls []int) {
	sls[0] = 100
}

func main() {
	a := [3]int{89, 90, 91}
	modify1(&a)
	fmt.Println(a)

	modify2(a[:])
	fmt.Println(a)
}
