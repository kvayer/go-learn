package day03

import "fmt"

func main() {
	var numbers []int

	numbers = append(numbers, 10)
	fmt.Printf("slice: %v len: %d cap: %d\n", numbers, len(numbers), cap(numbers))
	numbers = append(numbers, 20)
	fmt.Printf("slice: %v len: %d cap: %d\n", numbers, len(numbers), cap(numbers))
	numbers = append(numbers, 30)
	fmt.Printf("slice: %v len: %d cap: %d\n", numbers, len(numbers), cap(numbers))
	numbers = append(numbers, 40)
	fmt.Printf("slice: %v len: %d cap: %d\n", numbers, len(numbers), cap(numbers))
	numbers = append(numbers, 50)
	fmt.Printf("slice: %v len: %d cap: %d\n", numbers, len(numbers), cap(numbers))
}