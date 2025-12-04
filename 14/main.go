package main

import (
	"fmt"
	"reflect"
)

func DefineType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan")
	default: // проверка типа канала
		if isChannel(i) {
			fmt.Printf("channel %T\n", i)
		} else {
			fmt.Printf("don't know this %T type\n", i)
		}
	}
}

func isChannel(i interface{}) bool {
	t := reflect.TypeOf(i)
	return t != nil && t.Kind() == reflect.Chan
}

func main() {
	var i interface{}
	i = true
	DefineType(i)
	i = 5
	DefineType(i)
	i = "hello"
	DefineType(i)
	i = make(chan int)
	DefineType(i)
}
