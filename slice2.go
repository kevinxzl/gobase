package main

import (
	"fmt"
)

func main() {
	data := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := data[2:5]
	fmt.Println("array before", data)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after ", data)
}
