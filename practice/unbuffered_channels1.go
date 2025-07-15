package main


import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- time.Now().Format("03:04:05") // Blocks until reciever located
	}()

	// Note: reciever waits for all goroutines to finish basically waiting to maybe recieve something. If not, deadlock
	reciever := <- ch // Since main goroutine is faster than the goroutine that diverged, reciever is reached in time for channel to use
	fmt.Println(reciever)
}
