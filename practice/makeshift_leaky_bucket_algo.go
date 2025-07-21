package main

import (
	"fmt"
	"log"
	"time"
)

// Should try a mathmatical approach with time module later
type RateLimiter struct {
	token    chan struct{}
	leakRate time.Ticker
}

func NewRateLimiter(rateLimit int) *RateLimiter {
	rl := RateLimiter{
		token:    make(chan struct{}, rateLimit),
		leakRate: *time.NewTicker(1 * time.Second),
	}

	return &rl
}

func (rl *RateLimiter) Start() {
	go func() {
		for {
			select {
			case <-rl.leakRate.C:
				select {
				case <-rl.token:
					log.Println("Recieved token")
				default:
					log.Println("---- ---- ---- ----")
				}
			}
		}
	}()
}

func main() {
	rl := NewRateLimiter(10)

	rl.Start()

	for {
		time.Sleep(200 * time.Millisecond)
		select {
		case rl.token <- struct{}{}:
		default:
			fmt.Println("BUCKET FULL")
		}
	}

}
