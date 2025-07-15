package main

import (
	"fmt"
	"time"
)

func main() {
	c0 := time.After(7 * time.Second)
	c1 := make(chan string)
	c2 := make(chan string)

	go sendFast(c1)
	go sendSlow(c2)

Listener:
	for {
		select {
		case done := <-c0:
			fmt.Println("Request has timed out at:", done.Format("04:05"))
			break Listener
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}
}

func sendFast(ch chan string) {
	for i := 0; i <= 10; i++ {
		ch <- "Message from 1 at " + time.Now().Format("04:05")
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func sendSlow(ch chan string) {
	for i := 0; i <= 10; i++ {
		ch <- "\t\t\t\t\tMessage from 2 at " + time.Now().Format("04:05")
		time.Sleep(2 * time.Second)
	}
	close(ch)
}
