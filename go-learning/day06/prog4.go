package day06

import "fmt"

func printType(x any) {
	switch v := x.(type) {
	case int:
		fmt.Println("Int:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Bool:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	printType(42)
	printType("hello")
	printType(true)
	printType(3.14)
}