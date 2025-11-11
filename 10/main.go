package main

import "fmt"

func round(num float64) int {
	return int(num) / 10 * 10
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)
	for _, num := range arr {
		groups[round(num)] = append(groups[round(num)], num)
	}
	for k, num := range groups {
		fmt.Printf("%d: %v\n", k, num)
	}

}
