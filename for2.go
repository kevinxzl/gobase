package main

import "fmt"

func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	fmt.Printf("a = %d\n", a)

	/* for 循环 */
	fmt.Println("Test Case 1")
	for a := 0; a < 10; a++ {
		fmt.Printf("%d---", a)
	}
	fmt.Println()

	fmt.Println("Test Case 2")
	for a < b {
		a++
		fmt.Printf("%d---", a)
	}
	fmt.Println()

	fmt.Println("Test Case 3")
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	fmt.Println("Test Case 4")
	for v := range numbers {
		fmt.Printf("%d---", v)
	}
	fmt.Println()

	fmt.Println("Test Case 5")
	for _, v := range numbers {
		fmt.Printf("%d---", v)
	}
	fmt.Println()
}
