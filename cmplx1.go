package main
import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	fmt.Println("Test Case 1:")
	c := 3 + 4i
	r := cmplx.Abs(c)
	fmt.Println("The r =", r)
}

func main(){

	euler();
	//e^ipi + 1 = 0
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi) + 1)
}