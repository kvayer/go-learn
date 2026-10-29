package day02

import "fmt"

func printEvenNumbers(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	printEvenNumbers(n)
}