package basics 

import "fmt"

func main() {
	increment(10)
}

func increment(i int) {
	defer fmt.Println("first i value:", i)
	defer fmt.Println("Defering statement")
	i++
	fmt.Println("i is equal to:", i)
}
