package main

import (
	"context"
	"fmt"
	"time"
)

func stoppedByContextTimeout(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stopped by context timeout %d\n", id)
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Worker %d working...\n", id)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 0; i < 3; i++ {
		go stoppedByContextTimeout(ctx, i)
	}
	<-ctx.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("All workers finished - end program")

}
