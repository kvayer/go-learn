package day05

type Person struct {
	Name      string
	Age       int
	Height    float64
	IsStudent bool
}

func birthday(person Person) {
	person.Age += 1
}

func main() {
	person := Person{
		Name: "Mike",
		Age: 22,
		Height: 168,
		IsStudent: true,
	}
	fmt.Println(person.Age)
	birthday(person)
	fmt.Println(person.Age)
}