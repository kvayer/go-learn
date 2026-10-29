package day03

import "fmt"

func main() {
	var arr [5]int

	for i := 1; i <= len(arr); i++ {
		arr[i-1] = i * 10
	}

	fmt.Println(arr)
	fmt.Println(arr[1])
	fmt.Println(arr[3])
	fmt.Println(arr[4])
	arr[2] = 999
	fmt.Println(arr)
}