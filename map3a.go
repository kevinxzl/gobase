package main

import (
	"fmt"
)

func test(s string) {
	for i, ch := range []rune(s) {
		fmt.Println(i, ch)
	}
}

func main() {
	test("abccd")

}
