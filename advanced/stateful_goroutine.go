package advanced

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	ch    chan int
	count int
}

func (w *StatefulWorker) start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("Current value:", w.count)
			}
		}
	}()
}

func (w *StatefulWorker) send(value int) {
	w.ch <- value
}

func main() {
	stWorker := StatefulWorker{
		ch:    make(chan int),
		count: 0,
	}

	stWorker.start()

	for i := range 5 {
		stWorker.send(i)
		time.Sleep(time.Second)
	}

	fmt.Println("Main value:", stWorker.count)
}
