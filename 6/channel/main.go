package main

import (
	"fmt"
	"time"
)

func stoppedByChannel(stop <-chan struct{}, id int) {
	for {
		select {
		case <-stop:
			fmt.Printf("Worker %d stopped\n", id)
			return
		default:
			fmt.Printf("Worker %d running\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ch := make(chan struct{})

	for i := 0; i < 3; i++ {
		go stoppedByChannel(ch, i)
	}
	time.Sleep(2 * time.Second)

	close(ch)

	time.Sleep(500 * time.Millisecond)
	fmt.Println("All workers finished - end program")
}
