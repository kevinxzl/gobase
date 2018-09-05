package main

import (
	"fmt"
	"math"
)

var (
    aa = 3
    ss = "kkk"
    bb = true
)

func variableZeroValue() {
    fmt.Println("Test Case 1")
    var a int
    var s string
    fmt.Printf("%d %q\n", a, s)
}

func variableInitValue(){
    fmt.Println("Test Case 2")
    var a, b, c int = 3, 4, 5
    var s string = "abc"
    fmt.Println(a, b, c, s)
}

func variableTypeDeduction() {
    fmt.Println("Test Case 3")
    var a, b, c, s = 3, 4, true, "ABC"
    fmt.Println(a, b, c, s ,s)
}

func variableShorter(){
    fmt.Println("Test Case 4")
    a, b, c, s := 3, 4, true, "defavvh"
    fmt.Println(a, b, c, s)
}

func triangle() {
    fmt.Println("Test Case 5")
    var a, b int = 3, 4
    var c int
    c = int(math.Sqrt(float64( a*a + b*b)))
    fmt.Println(c)
}

func consts() {
    fmt.Println("Test Case 6")
    const ( 
        filename = "abc.txt"
        a, b = 3, 4
    )
    var c int
    c = int(math.Sqrt( a*a + b*b))
    fmt.Println(filename, c)
}

func enums1() {
    fmt.Println("Test Case 7")
    const (
        cpp = 0
        java = 1
        python = 2
        golang = 3
    )
    fmt.Println(cpp, java, python, golang)
}

func enums2() {
    fmt.Println("Test Case 8")
    const (
        cpp = iota
        java 
        python 
        golang 
    )
    const (
        b = 1 << (10 * iota)
        kb
        mb
        gb
        tb
        pb
    )
    fmt.Println(cpp, java, python, golang)
    fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
    fmt.Println("hello, world")
    variableZeroValue()
    variableInitValue()
    variableTypeDeduction()
    variableShorter()
    fmt.Println("Test Case 5")
    fmt.Println(aa, ss, bb)
    triangle()
    consts()
    enums1()
    enums2()
}
