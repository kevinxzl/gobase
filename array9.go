package main

import (
	"fmt"
)

func printArray1(arr [5]int){
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int){
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5] int
	arr2 := [3]int {1,3,5}
	arr3 := [...]int{2,4,6, 8, 10}
	fmt.Println(arr1, arr2, arr3)

	var grid [4][5] int
	fmt.Println(grid)

	fmt.Println("arr3:")
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	fmt.Println("arr3: i v")
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	fmt.Println("arr3: _, v")
	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray1(arr1)
	//printArray1(arr2)//cannot use arr2 (type [3]int) as type [5]int in argument to printArray
	printArray1(arr3)

	printArray2(&arr1)
	//printArray2(&arr2)
	printArray2(&arr3)
	

}