package advanced 

import(
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Goroutines")
	go nums()
	go alpha()
	fmt.Println("Welcome back to main thread")

	time.Sleep(2 * time.Second)
}

func nums() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func alpha() {
	for _, letter := range "abcdef" {
		fmt.Println(string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}
