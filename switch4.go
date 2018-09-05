package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("1 x 的类型 :%T\n", i)
	case int:
		fmt.Printf("2 x 是 int 型\n")
	case float64:
		fmt.Printf("3 x 是 float64 型\n")
	case func(int) float64:
		fmt.Printf("4 x 是 func(int) 型\n")
	case bool, string:
		fmt.Printf("5 x 是 bool 或 string 型\n")
	default:
		fmt.Printf("6 未知型\n")
	}
}
