package day02

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide: division by 0")
	}
	return a/b, nil
}

func main() {
	var a, b int

	fmt.Print("Input a: ")
	fmt.Scan(&a)
	fmt.Print("Input b: ")
	fmt.Scan(&b)

	result, err := divide(a, b)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println("Result:", result)
}