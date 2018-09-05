package main

import "fmt"

func main() {
	var map1 map[string]string
	/* 创建集合 */
	map1 = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	map1["Beijing"] = "China"
	map1["France"] = "Paris"
	map1["Italy"] = "Rome"
	map1["Japan"] = "Tokyo"
	map1["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range map1 {
		fmt.Println("Capital of", country, "is the country", map1[country])
	}

	str := "Beijing"

	/* 查看元素在集合中是否存在 */
	country, ok := map1[str]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Printf("find the Capital: %s is %s\n", str, country)
	} else {
		fmt.Println("Capital of %s  is not present\n", str)
	}
}
