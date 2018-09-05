package main

import (
	"fmt"
)

func main() {
	//
	for i := 0; i <= 10; i++ {
		if i > 5 && i < 8 {
			break //loop is terminated if i > 5 && i < 8
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
