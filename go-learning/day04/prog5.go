package day04

import "fmt"

func main() {
	scores := map[string]int{
		"Mike":  0,
		"Mike2": 10,
		"Mike3": 20,
		"Mike4": 30,
	}
	fmt.Println(scores)
	delete(scores, "Mike2")
	fmt.Println(scores)

	value, ok := scores["Mike2"]

	if ok {
		fmt.Println("Mike2 exists. Value:", value)
	} else {
		fmt.Println("Mike2 doesn't exists.")
	}
}