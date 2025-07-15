package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	c0 := time.After(7 * time.Second)
	c1 := make(chan string)
	c2 := make(chan string)

	go message(c1, "Maidison", 1)
	go message(c2, "Dall", 3)
	Data:
	for {
		select {
		case done := <- c0:
			fmt.Println("request timed out:", done.Format("03:04:05"))
			break Data
		case msg1 := <- c1:
			log.Println(msg1)
		case msg2 := <- c2:
			log.Println("\t\t\t\t", msg2)
		default:
			fmt.Println("waiting for data")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func message(ch chan string, name string, multi int){
	for {
		ch <- "This is a message from: " + name
		time.Sleep(time.Duration(multi) * time.Second)
	}
}
