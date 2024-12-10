package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

func NewHuman(name string, surname string, age int) Human {
	return Human{
		Name:    name,
		Surname: surname,
		Age:     age,
	}
}

func (h *Human) String() string {
	return fmt.Sprintf("%s %s, Age: %d", h.Name, h.Surname, h.Age)
}

type Action struct {
	Human
}

func main() {
	worker := NewHuman("John", "Jonson", 21)

	workerAction := Action{worker}

	fmt.Println(workerAction.String())
}
