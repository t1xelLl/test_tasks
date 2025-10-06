package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// без ввода в командной строки не заработает код
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <number_of_workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers")
		return
	}

	dataChan := make(chan string)

	var wg sync.WaitGroup
	// создаем воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}
	// пишем в канал
	go func() {
		counter := 0
		for {
			data := fmt.Sprintf("Data %d", counter)
			dataChan <- data
			counter++
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)+50))
		}
	}()

	wg.Wait()
}

// читаем из канала
func worker(id int, dataChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range dataChan {
		fmt.Printf("Worker %d: %s\n", id, data)
	}
}
