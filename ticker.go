package main

import (
	"fmt"
	"time"
)

func main() {
	for now := range time.Tick(2 * time.Second) {
		fmt.Println(now)
	}
}
