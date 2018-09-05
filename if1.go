package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10
	fmt.Println("Test Case 1:")
	fmt.Printf("a 的值为 : %d\n", a)

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	}

	fmt.Println("Test Case 2:")
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

}
