package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	numWorks := flag.Int("w", 1, "number of workers")
	timeout := flag.Int("t", 10, "timeout in seconds")
	flag.Parse()

	ch := make(chan int)

	wg := sync.WaitGroup{}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ctx, cancel := context.WithTimeout(ctx, time.Duration(*timeout)*time.Second)
	defer cancel()

	for i := 0; i < *numWorks; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)

	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- rand.Intn(100)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	fmt.Print("All workers finished - end program")
}

func worker(ctx context.Context, n int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped by context\n", n)
			return
		case data, ok := <-ch:
			if !ok {
				fmt.Printf("worker %d stopped: channel closed\n", n)
				return
			}
			fmt.Printf("worker %d processed: %d\n", n, data)
		}
	}
}
