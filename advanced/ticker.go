package advanced 

import (
	"fmt"
	"time"
)

// Accidently moved timer into ticker.go. oops lol

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 is dones")

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 finished")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Stopped timer 2 in main")
	}

	time.Sleep(2 * time.Second)
}
