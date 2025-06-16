package basics 

import "fmt"

func main() {
	applyTimes2 := multiplyByFactor(2)

	fmt.Println("5 * 2 =", applyTimes2(5))
}


func multiplyByFactor(factor int) func(int) int {
	return func(x int) int {
		return x * factor
		}
}
