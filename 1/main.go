package main

import "fmt"

type Human struct {
	name    string
	age     int
	country string
}

type Action struct {
	Human
	actionType string
}

func (h *Human) New() {
	fmt.Printf("Hello, my name is %s. I'm %d years old. I from %s.\n", h.name, h.age, h.country)
}

func (a *Action) Something() {
	fmt.Printf(a.actionType)
}

func main() {
	a := &Action{
		Human: Human{
			name:    "Tim",
			age:     20,
			country: "Russian",
		},
		actionType: "something",
	}
	fmt.Println(a.age)
	a.New()       // метод Human
	a.Something() // метод Action

}
