package advanced 

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	value int64
}

func (ac *atomicCounter) increment() {
	atomic.AddInt64(&ac.value, 1)
}

func (ac *atomicCounter) getVal() int64 {
	return atomic.LoadInt64(&ac.value)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	count := &atomicCounter{}

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				count.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter: %d\n", count.getVal())
}
