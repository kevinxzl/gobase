package main

import (
	"fmt"
)

func rectProps(l, w float64) (float64, float64) {
	var area = l * w
	var perimeter = (l + w) * 2
	return area, perimeter
}
func main() {
	area, _ := rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f\n ", area)
}
