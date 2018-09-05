package main

import "fmt"

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(a))

	for _, v := range a {
		fmt.Printf("%.2f---", v)
	}
	fmt.Println()

	fmt.Println(a)

}
