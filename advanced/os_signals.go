package advanced

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	pid := os.Getpid()
	fmt.Println(pid)

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("Recieved sginal:", sig)
		fmt.Println("Graceful exit.")
		os.Exit(0)
	}()

	fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}
