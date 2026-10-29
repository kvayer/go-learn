package day05

import "fmt"

type Person struct {
	Name      string
	Age       int
	Height    float64
	IsStudent bool
}

func main() {
	human := Person{
		Name:      "Mike",
		Age:       22,
		Height:    168,
		IsStudent: true,
	}

	fmt.Println(human)
	fmt.Println(human.Name, human.Age)
	human.Age = 23
	fmt.Println(human)
}