package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go sum(c, 5, -19) // Comes last?
	go sum(c, 4, 2, 10) // Is first?
	x, y := <-c, <-c // x = 6, y = -14

	fmt.Println(x, y, "x+y=", x+y)
}

func sum(c chan int, nums ...int) { // Intrestingly, with channel, the return value goes to the channel instead of actually being returned
	sum := 0
	for _, i := range nums {
		time.Sleep(100 * time.Millisecond)
		sum += i
	}
	c <- sum
}

