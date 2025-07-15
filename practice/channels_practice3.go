package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ch := make(chan string)

	go data("Message", ch)

	for {
		message, isChannelOpen := <-ch
		if !isChannelOpen {
			fmt.Println("Channel has been closed")
			break
		}
		log.Println(message)
	}

	fmt.Println("Congrats to me! Im out of the loop!")
}

func data(msg string, ch chan string) {

	for i := 0; i <= 10; i++ {
		ch <- msg + " at " + time.Now().Format("04:05")
		time.Sleep(2 * time.Second)
	}

	closeChannel(ch)
}

func closeChannel(ch chan string) {
	close(ch)
}
