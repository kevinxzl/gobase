package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {
	/* 创建切片 */
	//numbers := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("1 numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("2 numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("3 numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("4 numbers[4:] ==", numbers[4:])

	fmt.Println("5")
	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	fmt.Println("6")
	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]

	printSlice(number2)

	fmt.Println("7")
	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}
