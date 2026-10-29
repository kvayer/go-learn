package day04

import "fmt"

func main() {

	guys := map[string]int {
		"Mike": 22,
		"Anna": 19,
		"John": 23,
	}

	fmt.Println(guys)
	fmt.Println(guys["Mike"])
	fmt.Println(guys["John"])
	guys["Mike"] = 23
	guys["Alex"] = 25
	fmt.Println(guys)

}