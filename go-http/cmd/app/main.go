package main

import (
	"net/http"
)

func main() {

	repo := NewMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	mux := NewRouter(handler)

	http.ListenAndServe(":8080", mux)
}

func NewRouter(handler *UserHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handler.Users)
	mux.HandleFunc("/users/{id}", handler.UsersByID)
	return mux
}
