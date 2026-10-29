package day02

import "fmt"

func main() {
	var age int
	var haveTicket bool
	var inBL bool
	var isVIP bool

	isAdult := age >= 18

	fmt.Print("Возраст: ")
	fmt.Scan(&age)

	fmt.Print("Есть билет: ")
	fmt.Scan(&haveTicket)

	fmt.Print("В бан-листе: ")
	fmt.Scan(&inBL)

	fmt.Print("VIP: ")
	fmt.Scan(&isVIP)

	if isAdult && isVIP && !inBL {
		fmt.Println("Доступ разрешён")
	} else if isAdult && haveTicket && !inBL{
		fmt.Println("Доступ разрешён")
	} else {
		fmt.Println("Доступ запрещён")
	}
}