package main

import (
	"fmt"
	"packages-demo/internal/calculator"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(calculator.Sum(50, 50))
	fmt.Println(uuid.New().String())
}

