package main

import "fmt"

func main() {
	var a int = 20 /* 声明实际变量 */
	var p *int     /* 声明指针变量 */

	p = &a /* 指针变量的存储地址 */
	fmt.Println("Test Case 1")
	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("p 变量的存储地址: %x\n", p)

	/* 使用指针访问值 */
	fmt.Printf("*p 变量的值: %d\n", *p)

	fmt.Println("Test Case 2")
	*p = 50
	fmt.Println(a, *p)

	fmt.Println("Test Case 3")
	b3 := 255
	a3 := &b3
	fmt.Println("address of b3 is", a3)
	fmt.Println("value of b3 is", *a3)
}
