package basics 

import "fmt"

func main() {
	isPositive(5)

	isPositive(-3)
}

func isPositive(i int) {
	defer fmt.Printf("Value of i: %d\n", i)

	if i < 0 {
		panic("Value i is a negative number")
	}
	fmt.Printf("i is a positive number")
}
