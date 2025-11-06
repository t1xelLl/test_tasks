package main

import (
	"fmt"
)

const lenArray = 10

func main() {
	arr := make([]int, lenArray)
	// запись в слайс от 1 до N
	for i := 0; i < lenArray; i++ {
		arr[i] = i
	}
	ch1 := make(chan int)
	ch2 := make(chan int)

	// записываем значения в 1 канал
	go func() {
		defer close(ch1)
		for _, i := range arr {
			ch1 <- i
		}
	}()

	// читаем значения из 1 канала и записываем квадраты чисел во 2 канал
	go func() {
		defer close(ch2)
		for i := range ch1 {
			ch2 <- i * i
		}
	}()

	// вывод
	for i := range ch2 {
		fmt.Print(i, " ")
	}
}
