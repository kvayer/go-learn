package day02

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	var a, b int

	fmt.Print("Введите a: ")
	fmt.Scan(&a)
	fmt.Print("Введите b: ")
	fmt.Scan(&b)
	fmt.Printf("Сумма a и b: %d", sum(a,b))
}