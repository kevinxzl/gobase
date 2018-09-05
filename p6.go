package main

import "fmt"

const MAX int = 3

func main() {

	a := []int{10, 100, 200}

	fmt.Println("Test Case 1")
	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	//Go不支持指针算法。没有 p++ p-- 操作 ，故可以把数组的地址赋给一个数组指针
	fmt.Println("Test Case 2")
	var parr [MAX]*int
	for i := 0; i < MAX; i++ {
		parr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *parr[i])
	}
}
