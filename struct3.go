package main

import "fmt"

type Books struct {
}

func (s Books) String() string {
	return "my data"
}
func main() {
	fmt.Printf("%v\n", Books{})
}
