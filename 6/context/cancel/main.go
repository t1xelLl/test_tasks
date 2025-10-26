package main

import (
	"context"
	"fmt"
	"time"
)

func stoppedByContextWithCancel(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stopped by context: %d \n", id)
			return
		default:
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		go stoppedByContextWithCancel(ctx, i)
	}

	time.Sleep(2 * time.Second)

	cancel()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("All workers finished - end program")
}
