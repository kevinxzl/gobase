package main

import (
	"fmt"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap1(x, y int) (int, int) {
	return y, x
}

func print(arr []int) {
	for v := range arr {
		fmt.Println(v, "  ")
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			swap(&arr[i], &arr[j])
		}
	}
	i++
	swap(&arr[i], &arr[high])
	return i
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	data := [...]int{90, 83, 82, 67, -56, 44, 85, -34, 10, 1, 3, -45, 24, 0, 34, 78, 69, -1}
	arr := data[:]
	fmt.Println("Orginal array is:")
	fmt.Println(arr)
	high := len(arr) - 1
	//fmt.Println(high)
	quickSort(arr, 0, high)
	fmt.Println("After SORT array is:")
	fmt.Println(arr)

}
