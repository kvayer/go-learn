package day02

import "fmt"

func main() {
	var n = 100
	var summ int

	for i := 1; i <= n; i++ {
		summ += i
	}

	fmt.Println("Сумма:", summ)
}