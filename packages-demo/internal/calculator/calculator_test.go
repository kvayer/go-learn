package calculator

import (
	"errors"
	"testing"
)

func TestSum(t *testing.T) {

	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive and zero", 5, 0, 5},
		{"negative and positive", -10, 5, -5},
		{"negative and negative", -10, -10, -20},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual := Sum(test.a, test.b)
			assertEqual(t, actual, test.expected)
		})

	}
}

func assertEqual(t *testing.T, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected: %d, got: %d", expected, actual)
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		name           string
		a              int
		b              int
		expectedResult int
		wantError      error
	}{
		{name: "negative by positive", a: -20, b: 5, expectedResult: -4, wantError: nil},
		{name: "division by zero", a: 10, b: 0, expectedResult: 0, wantError: ErrDivisionByZero},
		{name: "positive by positive", a: 16, b: 4, expectedResult: 4, wantError: nil},
		{name: "negative by negative", a: -25, b: -5, expectedResult: 5, wantError: nil},
		{name: "identical numbers", a: 2345623, b: 2345623, expectedResult: 1, wantError: nil},
		{name: "zero by positive", a: 0, b: 2345623, expectedResult: 0, wantError: nil},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual, err := Divide(test.a, test.b)
			assertDivision(t, actual, test.expectedResult, err, test.wantError)
		})
	}
}

func assertDivision(t *testing.T, actual, expectedResult int, err, wantError error) {
	t.Helper()

	if !errors.Is(err, wantError) {
		t.Errorf("expected error: %v, got: %v", wantError, err)
		return
	}

	if wantError != nil {
		return
	}

	if actual != expectedResult {
		t.Errorf("expected result: %d, got: %d", expectedResult, actual)
	}
}
