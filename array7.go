package main

import "fmt"

// pass by value
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func getAverage(arr [5]int, size int) (sum int, avg float32) {
	for i := 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return sum, avg
}

func main() {
	nums := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", nums)
	changeLocal(nums) //num is passed by value
	fmt.Println("after passing to function ", nums)
	sum, avg := getAverage(nums, 5)
	fmt.Printf("The array sum=%d  avg=%.2f\n", sum, avg)
}
