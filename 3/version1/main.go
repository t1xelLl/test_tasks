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
	// определяем количество воркеров
	var numWorkers int = runtime.NumCPU()
	if len(os.Args) > 1 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil && n > 0 {
			numWorkers = n
		}
	}
	// создаем контекст для управления жизненным циклом горутин
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// канал для передачи данных
	ch := make(chan string)
	wg := sync.WaitGroup{}
	// запуск N воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	wg.Add(1)

	go writer(ctx, ch, &wg)

	// Gracefull showdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nReceived shutdown signal, shutting down gracefully")
	cancel()
	wg.Wait()

	close(ch)
	fmt.Println("end")
}

// пишет в канал
func writer(ctx context.Context, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 1
	for {
		select {
		case <-ctx.Done(): // проверка отмены контекста
			fmt.Println("writer is done")
			return
		default: // создаем запись
			data := fmt.Sprintf("Message %d (time: %v)", counter, time.Now().Format("15:04:05"))

			select {
			case ch <- data: // пытаемся отправить данные
				counter++
			case <-ctx.Done():
				return // проверка отмены во время отправки
			}

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)+50))
		}
	}
}

// читает из канала
func worker(ctx context.Context, id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // отмена контекста
			fmt.Printf("worker %d finish work\n", id)
			return
		case data, ok := <-ch: // получаем данные из канала
			if !ok {
				fmt.Printf("worker %d close channel\n", id)
				return
			}
			fmt.Printf("worker %d received data: %v\n", id, data)
		}
	}
}
