package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{
		"1steve": 100,
		"2jamie": 200,
		"3tim":   300,
	}
	map1["4mike"] = 400
	fmt.Println("map1", map1)
	map2 := map1
	map2["4mike"] = 480
	map1["3tim"] = 380
	fmt.Println("changed", map1)
	fmt.Println("map2   ", map2)

	arr1 := [...]int{1, 2, 3, 4}
	arr2 := arr1
	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)
	fmt.Println("Changed arr1 and arr2 value")
	arr1[0] = 9
	arr2[0] = 8
	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)

}
