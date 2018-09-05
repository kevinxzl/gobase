package main

import "fmt"

func main() {
	/* 创建 map */
	//key: country   value: captial
	map1 := map[string]string{"1China": "Beijing", "2France": "Paris", "3Italy": "Rome", "4Japan": "Tokyo", "5India": "New Delhi"}

	fmt.Println("原始 map")

	/* 打印 map */
	for key := range map1 {
		fmt.Println("Capital of", key, "is", map1[key])
	}

	/* 删除元素 */
	delete(map1, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range map1 {
		fmt.Println("Capital of", country, "is", map1[country])
	}
}
