package gpt_assignments

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

type RateLimiter struct {
	signal   chan struct{}
	rate     time.Ticker
	tokens   int64
	capacity int64
}

func (rl *RateLimiter) FillBucket() {
	for {
		select {
		case <-rl.signal:
			if rl.capacity > rl.tokens {
				atomic.AddInt64(&rl.tokens, 1)
			}
		default:
		}
	}
}

func (rl *RateLimiter) MakeRequest() bool {
	if rl.tokens != 0 {
		atomic.AddInt64(&rl.tokens, -1)
		return true
	}
	return false
}

func (rl *RateLimiter) tick() {
	for {
		select {
		case <-rl.rate.C:
			rl.signal <- struct{}{}
		default:
		}
	}
}

func sendMessage() {
	log.Println("Message sent.")
}

// 7-18-25 GPT challanged me to make a token bucket algo, without having seen it before. Here was my attempt
// GPT's Comments:
/*
1.Busy Looping (FillBucket and tick)

You're using default in your select, which means these goroutines are spinning non-stop when there’s nothing to receive/send. That’s CPU waste and not how channels should be used.



2. Misuse of signal and rate Ticker

You're using a custom signal channel and a ticker in a weird handshake. But rate.C is already a channel—you don’t need the extra indirection.



3.Race condition in MakeRequest()

You check tokens != 0, then decrement. But two goroutines could pass that check before either subtracts, overshooting the capacity.



4. Poor test flow

You start the message loop after 10 seconds, but meanwhile the bucket isn't filling tokens—makes it look like nothing’s working.
*/
func main() {
	newRate := RateLimiter{
		signal:   make(chan struct{}),
		rate:     *time.NewTicker(1 * time.Second),
		tokens:   0,
		capacity: 25,
	}

	go newRate.tick()
	go newRate.FillBucket()

	messageRate := time.NewTicker(500 * time.Millisecond)
	waitForBucket := time.NewTimer(10 * time.Second)
	<-waitForBucket.C
	fmt.Println("Begin messages:")

	for {
		select {
		case <-messageRate.C:
			if newRate.MakeRequest() {
				sendMessage()
			} else {
				fmt.Println("Over rate")
			}
		}
	}
}
