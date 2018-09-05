package main

import "fmt"

func getSequence() func() int {
	i := 0
	fmt.Println("f1:", i)
	return func() int {
		i += 1
		fmt.Println("f2:", i)
		return i
	}
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println("1---", nextNumber())
	fmt.Println("2---", nextNumber())
	fmt.Println("3---", nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println("4--", nextNumber1())
	fmt.Println("5--", nextNumber1())
}
