package main

import "fmt"

func SetBitTo1(num int64, pos int) int64 {
	return num | (1 << pos)
}

func SetBitTo0(num int64, pos int) int64 {
	return num &^ (1 << pos)

}

func main() {
	fmt.Println(SetBitTo0(5, 0)) // 5 -> 4
	fmt.Println(SetBitTo1(5, 1)) // 5 -> 7
}
