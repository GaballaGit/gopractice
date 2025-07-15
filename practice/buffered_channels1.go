package main

import "fmt"

func main() {

	ch := make(chan string, 10)

	ch <- "Deadlocks?"
	ch <- "I"
	ch <- "Haven't"
	ch <- "Heard"
	ch <- "Of"
	ch <- "Them!"

	close(ch)

	for msg := range ch {
		fmt.Println(msg)
	}

}
