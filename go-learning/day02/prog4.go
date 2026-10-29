package day02

import "fmt"

func main() {
	var command int

	fmt.Print("Введите команду: ")
	fmt.Scan(&command)

	switch command {
		case 1:
		fmt.Println("Добавить")
		case 2:
		fmt.Println("Найти")
		case 3:
		fmt.Println("Удалить")
		case 4:
		fmt.Println("Выйти")
		default:
			fmt.Println("Команда не распознана")
	}
}