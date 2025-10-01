package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{2, 4, 6, 8}
	fmt.Println(SolutionWithWaitGroup(slice))
	fmt.Println(SolutionWithChannel(slice))

}

func SolutionWithWaitGroup(sl []int) []int {
	result := make([]int, len(sl))
	var wg sync.WaitGroup
	for i, num := range sl {
		wg.Add(1)
		go func(i int, num int) {
			defer wg.Done()
			result[i] = num * num
		}(i, num)
	}
	wg.Wait()
	return result
}

func SolutionWithChannel(sl []int) []int {
	result := make([]int, len(sl))

	type orderedResult struct {
		index int
		value int
	}

	ch := make(chan orderedResult, len(sl))

	for i, v := range sl {
		go func(idx, num int) {
			ch <- orderedResult{idx, num * num}
		}(i, v)
	}

	for i := 0; i < len(sl); i++ {
		res := <-ch
		result[res.index] = res.value

	}
	close(ch)
	return result
}
