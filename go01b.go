package main
import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)
	for i:=0; i < 10; i++ {
		go Go(c,i)
	}
	<-c
}

func Go(c chan bool, index int) {
	sum := 1;
  for i:=0; i < 100000000; i++ {
		sum += i;
	}
	fmt.Println(index, sum)

	if index == 9 {
		c <- true
	}
}
