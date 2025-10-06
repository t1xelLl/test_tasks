package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	var numWorkers int = runtime.NumCPU()
	if len(os.Args) > 1 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil && n > 0 {
			numWorkers = n
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan string)
	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	wg.Add(1)

	go writer(ctx, ch, &wg)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nReceived shutdown signal, shutting down gracefully")
	cancel()
	wg.Wait()

	close(ch)
	fmt.Println("end")
}

func writer(ctx context.Context, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writer is done")
			return
		default:
			data := fmt.Sprintf("Message %d (time: %v)", counter, time.Now().Format("15:04:05"))

			select {
			case ch <- data:
				counter++
			case <-ctx.Done():
				return
			}

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)+50))
		}
	}
}

func worker(ctx context.Context, id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d finish work\n", id)
			return
		case data, ok := <-ch:
			if !ok {
				fmt.Printf("worker %d close channel\n", id)
				return
			}
			fmt.Printf("worker %d received data: %v\n", id, data)
		}
	}
}
