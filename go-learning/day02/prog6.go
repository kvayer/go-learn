package day02

import "fmt"

func main() {
	var n int

	fmt.Print("Введиче число: ")
	fmt.Scan(&n)

	for i:= 1; i <= n; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}