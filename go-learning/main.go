package main

import "fmt"

func main() {
	s := "привет"
	sRune := []rune(s)
	fmt.Println(s[0])
	sRune[0] = 'П'
	fmt.Println(s[0], sRune[0])
	s = "hello"
	sRune = []rune(s)
	fmt.Println(s[0])
	sRune[0] = 'H'
	fmt.Println(s[0], sRune[0])
}