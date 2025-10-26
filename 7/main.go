package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	m := make(map[rune]int) // мапа для подсчета количества элементов
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	str := "Hello World!😀"
	runes := []rune(str)

	ch := make(chan rune, len(runes))

	wg.Add(1)
	// пишем в канал руны
	go func() {
		defer wg.Done()
		defer close(ch)
		for _, v := range runes {
			ch <- v
		}
	}()

	numWorkers := min(len(runes), runtime.NumCPU())
	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			for v := range ch {
				// с помощью мьютекса захватываем мапу блокируем доступ для других горутин
				mutex.Lock()
				m[v]++         // если нет ключа создастся ячейка с нулевым значением и мы сразу к ней прибавляем
				mutex.Unlock() // разблокировали
			}
		}()
	}
	wg.Wait()
	for k, v := range m {
		fmt.Printf("Rune: %c %d times\n", k, v)
	}

}
