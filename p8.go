package main

import "fmt"

func swap1(x, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
func swap2(x, y *int) {
	*x, *y = *y, *x
}

func swap3(x, y int) (int, int) {
	return y, x
}

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Println("swap 1")
	fmt.Println(a, b)
	swap1(&a, &b)
	fmt.Println(a, b)

	fmt.Println("swap 2")
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)

	fmt.Println("swap 3")
	fmt.Println(a, b)
	a, b = swap3(a, b)
	fmt.Println(a, b)

}
