package advanced 

import (
	"fmt"
	"sync"
)


type Counter struct{
	mu sync.Mutex
	value int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &Counter{}


	for range numGoroutines {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.getValue())
}
