package main

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	var numWorkers = runtime.NumCPU()

	// Создаем контекст, который автоматически отменяется при получении сигналов завершения
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	// используем для ожидания завершения всех горутин
	wg := sync.WaitGroup{}

	for i := range numWorkers {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done(): // этот case сработает когда контекст получить SIGINT/SIGTERM
					fmt.Printf("worker %d done\n", id)
					return
				default: // имитация полезной работы
					time.Sleep(3 * time.Second)
					fmt.Printf("worker %d did work\n", id)
				}
			}
		}(i)
	}

	// блокируемся пока счетчик WaitGroup не станет 0
	wg.Wait()

}
