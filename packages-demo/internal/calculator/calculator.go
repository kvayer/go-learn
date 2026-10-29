package calculator

import (
	"errors"
	"fmt"
)

var ErrDivisionByZero = errors.New("division by zero")

func Sum(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide numbers: %w", ErrDivisionByZero)
	}
	return a / b, nil
}
