package day03

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	a := numbers[1:4]

	fmt.Println(a, len(a), cap(a))
}