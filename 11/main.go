package main

import "fmt"

func Intersect[T comparable](sl1, sl2 []T) []T {
	// выбор меньшего слайса
	if len(sl1) < len(sl2) {
		sl1, sl2 = sl2, sl1
	}
	// создание set
	set := make(map[T]struct{})

	// заполнение set значениями меньшего слайса для проверок
	for _, v := range sl1 {
		set[v] = struct{}{}
	}

	res := make([]T, 0)

	// проход по 2-му слайсу и добавление в результирующий слайс значение которые есть в обоих слайсах
	for _, v := range sl2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	sl1 := []int{1, 2, 3}
	sl2 := []int{2, 3, 4}
	fmt.Println(Intersect(sl1, sl2))
}
