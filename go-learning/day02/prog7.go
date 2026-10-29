package day02

import "fmt"

func main() {

	var n int

	fmt.Print("Введите N: ")
	fmt.Scan(&n)

	for i:= 2; i <= n; i += 2 {
		fmt.Println(i)
	}
}