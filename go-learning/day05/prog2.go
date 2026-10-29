package day05

package day05

import "fmt"

type Person struct {
	Name      string
	Age       int
	Height    float64
	IsStudent bool
}

func main() {
	persons := []Person{
		Person{
			Name:      "Mike",
			Age:       22,
			Height:    168,
			IsStudent: true,
		},
		Person{
			Name:      "Nikita",
			Age:       22,
			Height:    198,
			IsStudent: true,
		},
		Person{
			Name:      "Dasha",
			Age:       33,
			Height:    160,
			IsStudent: false,
		},
	}

	for person := range persons {
		fmt.Printf("Имя: %s\nВозраст: %d\n\n", persons[person].Name, persons[person].Age)
	}
}
