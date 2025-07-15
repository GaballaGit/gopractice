package main

import(
	"time"
	"fmt"
)

func main() {
	go say("world") // Note: Since it branches off, main goroutine continues, unblocked by say("world")
	say("hello")
}

func say(text string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(text)
	}
}

