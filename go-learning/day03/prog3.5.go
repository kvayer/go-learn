package day03

import "fmt"

func main() {
	numbers := []int{5, 10, 15, 20}

	for index, value := range numbers {
		fmt.Printf("index=%d value=%d\n", index, value)
	}
	for _, value := range numbers {
		fmt.Printf("value=%d\n", value)
	}
}