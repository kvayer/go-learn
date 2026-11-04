package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {

	repo := NewMemoryUserRepository()
	service := NewUserService(repo)
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	handler := NewUserHandler(service, logger)

	mux := NewRouter(handler)

	http.ListenAndServe(":8080", mux)
}

func NewRouter(handler *UserHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handler.Users)
	mux.HandleFunc("/users/{id}", handler.UsersByID)
	return mux
}
