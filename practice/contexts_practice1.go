package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancle := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancle()


	ch := make(chan string)
	nameless := []string{"March 7th", "Dan Heng", "Trailblazer"}

	for _, name := range nameless {
		go runMission(name, ch)
	}

	Execute:
	for {
		select {
		case success := <- ch:
		fmt.Println(success)
		case <-ctx.Done():
		fmt.Println("Time has closed")
		break Execute
		}
	}
}

func runMission(name string, ch chan string) {
	time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
	ch <- name + " has completed their mission"
}

