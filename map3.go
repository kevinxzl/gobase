package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	key, ok := m["b"]
	fmt.Println(key, ok)
	key, ok = m["a"]
	fmt.Println(key, ok)
}
