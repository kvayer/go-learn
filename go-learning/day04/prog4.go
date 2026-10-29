package day04

import "fmt"

func main() {
	scores := map[string]int{
		"Mike": 0,
	}

	value, ok := scores["Mike"]
	if ok {
		fmt.Println("Mike is exists. Value:", value)
	}	else { 
		fmt.Println("Mike doesn't exists.")
	}
	value, ok = scores["Unknown"]
	if ok {
		fmt.Println("Unknown is exists. Value:", value)
	}	else { 
		fmt.Println("Unknown doesn't exists.")
	}
}