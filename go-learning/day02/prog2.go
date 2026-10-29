package day02
import "fmt"
// 0–12  → ребёнок
// 13–17 → подросток
// 18–64 → взрослый
// 65+   → пожилой
func main() {
	var age int
	fmt.Println("Введите возраст:")
	fmt.Scan(&age)
	if age > 0 && age < 13 {
		fmt.Println("Вы ребёнок")
	} else if age < 18 {
		fmt.Println("Вы подросток")
	} else if age < 65 {
		fmt.Println("Вы взрослый")
	} else {
		fmt.Println("Вы пожилой")
	} 
}