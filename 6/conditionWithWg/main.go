package main

import (
	"fmt"
	"sync"
	"time"
)

func stoppedByConditionWithWg(wg *sync.WaitGroup, stop *bool, id int) {
	defer wg.Done()

	for !*stop {
		fmt.Printf("Worker %d working...\n", id)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("Worker %d stopped\n", id)
}

func main() {
	wg := new(sync.WaitGroup)
	stop := false

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go stoppedByConditionWithWg(wg, &stop, i)
	}

	time.Sleep(2 * time.Second)
	stop = true
	wg.Wait()
	fmt.Println("All workers finished - end program")
}
