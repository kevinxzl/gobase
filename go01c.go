package main
import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i:=0; i < 10; i++ {
		go Go(c,i)
	}
	
	for i := 0; i < 10; i++ {
		<-c
	}

}

func Go(c chan bool, index int) {
	sum := 1;
  for i:=0; i < 100000000; i++ {
		sum += i;
	}
	fmt.Println(index, sum)
  
	c <- true
}
