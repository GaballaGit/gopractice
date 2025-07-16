package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	ch := make(chan string)
	var wg sync.WaitGroup

	reporters := []string{"Micha", "Alan", "Venus"}

	for _, i := range reporters {
		wg.Add(1)
		go report(i, ch, ctx, &wg)
	}

COMS:
	for {
		select {
		case <-ctx.Done():
			log.Println("Communications cut off.")
			break COMS
		case msg := <-ch:
			fmt.Println(msg)
		default:
			fmt.Println("waiting for response.")
			time.Sleep(500 * time.Millisecond)
		}
	}
	wg.Wait()
	close(ch)
	fmt.Println("TERMINATED")
}

func report(name string, ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s CONNECTION LOST\n", name)
			return
		case <-time.After(time.Duration(rand.Intn(3)+1) * time.Second):
			ch <- time.Now().Format("03:04:05") + ": " + name + " has reported data."
		}
	}
}
