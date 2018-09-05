package main
import (
	"io/ioutil"
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func grade(score int) string {
	g := ""
	switch{
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func convertToBin(n int) string {
	
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2;
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
//dead loop
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main(){
	fmt.Println("Test Case 1")
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println("Test Case 2")
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		//grade(200),
	)

	fmt.Println("Test Case 3")
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(54542542),
		convertToBin(0))

		fmt.Println("Test Case 4")
		printFile("abc.txt")
		fmt.Println("Test Case 5---dead loop")
		//forever()
}