package day07

import "fmt"

func main() {
	s := "hello"
	fmt.Println(s)
	byteH := "H"
	sByte := []byte(s)
	sByte[0] = byteH[0]
	fmt.Println(string(sByte))
}