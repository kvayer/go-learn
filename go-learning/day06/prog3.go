package day06

import "fmt"

type Greeter interface {
	Greet() string
}

type Robot struct {
	Model string
}

type Person struct {
	Age  int
	Name string
}

func (r Robot) Greet() string {
	return "Gr-ee-ti-ngs!"
}

func (p Person) Greet() string {
	return "Greetings!"
}

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	person := Person{Name: "Mike", Age: 22}
	sayHello(person)

	robot := Robot{Model: "Tesla"}
	sayHello(robot)
}
