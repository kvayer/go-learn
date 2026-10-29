package day06

import (
	"fmt"
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

func (p Person) Info() string {
	return fmt.Sprintf("%s, %d years old",p.Name, p.Age)
}

func main() {
	person := Person{Name: "Mike", Age: 22}
	fmt.Println(person.Name, person.Age)
	person.Rename("Michael")
	fmt.Println(person.Name, person.Age)
	person.Birthday()
}
