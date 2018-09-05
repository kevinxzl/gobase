package main

import (
	"fmt"
)

func main() {
	a := [6]int{76, 77, 78, 79, 80, 90}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(a)
	fmt.Println(b)
}
