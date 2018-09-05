package main

import (
	"fmt"
	"time"
)

var a[10] int
func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i] ++
		   }
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}

