package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.Mutex
	requests int64
	window   time.Ticker
}

func (rl *RateLimiter) Start(requests int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	go func() {
		for {
			select {
			case <-rl.window.C:
				rl.requests = int64(requests)
				fmt.Println("Requests refresed")
			}
		}
	}()
}

func (rl *RateLimiter) MakeRequest() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	if rl.requests <= 0 {
		fmt.Println("Rate limited")
		return false
	} else {
		rl.requests -= 1
		return true
	}
}

func main() {
	requestLimit := 10

	rl := RateLimiter{
		mu:       sync.Mutex{},
		requests: int64(requestLimit),
		window:   *time.NewTicker(5 * time.Second),
	}

	rl.Start(requestLimit)

	for {
		if rl.MakeRequest() {
			log.Println("message sent")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
