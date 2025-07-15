package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}


	wg.Add(4)
	ch := make(chan int)

	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	go reader(3, ch, &wg)
	go reader(4, ch, &wg)


	for i := 0; i <= 100; i++ {
		ch <- i
	}

	x, y, z, w := <-ch, <-ch, <-ch, <-ch
	fmt.Println(x,y,z,w)
	close(ch)

	wg.Wait()

}

func reader(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		val, ok := <-ch 
		if !ok {
			fmt.Println("no more channel input")
			ch <- id
			return
		}
		fmt.Printf("Channel %d has recieved %d\n", id, val)
	}

}
