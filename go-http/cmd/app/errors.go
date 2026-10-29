package main

import "errors"

var ErrInvalidUserData = errors.New("invalid user data")
var ErrUserNotFound = errors.New("user not found")
