package day02

import "fmt"

func main() {
	var n int

	fmt.Print("Введите N: ")
	fmt.Scan(&n)

	for i:= 1; i <= n; i++ {
		fmt.Println(i)
	}
}