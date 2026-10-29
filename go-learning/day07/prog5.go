package day07

import "fmt"

func main() {
	s := "привет"
	sRune := []rune(s)
	sRune[0] = 'П'
	fmt.Println(string(sRune))
}