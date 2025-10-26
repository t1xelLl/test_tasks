package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	start := time.Now()

	go func() {
		for {
			fmt.Println("Goroutine is running")
			time.Sleep(500 * time.Millisecond)

			if time.Since(start).Seconds() >= 2 {
				fmt.Println("Goroutine is stopped")
				runtime.Goexit()
			}
		}
	}()
	time.Sleep(3 * time.Second)
}
