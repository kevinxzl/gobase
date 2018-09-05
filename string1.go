package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Yes! 这是中国！" //UTF-8
	fmt.Println("s1 len=", len(s1), "content=", s1)
	fmt.Printf("%s\n", []byte(s1))

	for _, b := range []byte(s1) {
		fmt.Printf("%X  ", b)
	}
	fmt.Println()
	for i, ch := range s1 { // ch is a rune
		fmt.Printf("(%d  %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s1))
	fmt.Println()

	bytes := []byte(s1)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c---", ch)
	}
	fmt.Println()
	for i, ch := range []rune(s1) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
