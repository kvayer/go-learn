package day02

import "fmt"

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	var even int
	var evenSum int

	for i:= 1; i <= n; i++ {
		if i%2 == 0 {
			even++
			evenSum += i
		}
	}
	fmt.Printf("Количество чётных: %d\nСумма чётных: %d", even, evenSum)
}