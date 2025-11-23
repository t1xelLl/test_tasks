package main

import "fmt"

func NewSet[T comparable](sl []T) []T {
	set := make(map[T]struct{})

	for _, v := range sl {
		set[v] = struct{}{}
	}

	res := make([]T, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	return res
}

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(NewSet(sl))
}
