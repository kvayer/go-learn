package day02

import "fmt"

func main() {
	var age int

	fmt.Println("Введите возраст:")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Вы совершеннолетний.")
	} else {
		fmt.Println("Вы несовершеннолетний.")
	}
}