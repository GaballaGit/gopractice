package basics

import "fmt"

func main() {
	fmt.Printf("%d\n", sums(5, 4, 2, 7, 10, 2, 61, 100))
}

func sums(nums ...int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum
}
