package main

import "errors"

type FieldError struct {
	Field string
}

func (fe *FieldError) Error() string {
	return "bad field: " + fe.Field
}

var ErrInvalidUserData = errors.New("invalid user data")
var ErrUserNotFound = errors.New("user not found")
