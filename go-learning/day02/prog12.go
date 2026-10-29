package day02

import "fmt"

func calculate(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	var a, b, sum, difference int

	fmt.Print("Введите a: ")
	fmt.Scan(&a)
	fmt.Print("Введите b: ")
	fmt.Scan(&b)

	sum, difference = calculate(a, b)

	fmt.Printf("Сумма: %d\nРазница: %d", sum, difference)
}