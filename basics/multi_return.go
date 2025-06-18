package basics 

import "fmt"

func main() {
	q, r := div(37, 8)
	fmt.Printf("Dividing 37 with 8 we get %d and a remainder of %d\n", q, r) 
}

func div(x int, y int) (int, int) {
	quotiant := x / y
	remainder := x % y
	return quotiant, remainder
}
