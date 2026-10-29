package day06

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Birthday() {
	p.Age++
}

func (p *Person) Rename(newName string) {
	p.Name = newName
}

func main() {
	person := Person{Name: "Mike", Age: 22}
	fmt.Println(person.Name, person.Age)
	person.Rename("Michael")
	fmt.Println(person.Name, person.Age)
	person.Birthday()
}
