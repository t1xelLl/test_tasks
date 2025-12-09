package main

import "fmt"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	return string(make([]byte, size))
}
