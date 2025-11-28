package main

func main() {
	a, b := 1, 2
	a = a ^ b
	b = a ^ b
	a = a ^ b
	println(a, b)
}
