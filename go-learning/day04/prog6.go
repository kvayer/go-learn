package day04

import "fmt"

func plus(m map[string]int, key string) {
	m[key] += 1
}

func main() {
	m := make(map[string]int)
	m["Mike"] = 22
	fmt.Println(m)
	plus(m, "Mike")
	fmt.Println(m)
}