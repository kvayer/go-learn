package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

const waitingTime = 2 * time.Second

type UserHandler struct {
	service *UserService
	logger  *slog.Logger
}

func NewUserHandler(service *UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	switch r.Method {
	case http.MethodPost:
		var input CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		user, err := h.service.CreateUser(ctx, input.Name, input.Age)
		if err != nil {
			var fe *FieldError
			if errors.As(err, &fe) {
				http.Error(w, fmt.Sprintf("400 Bad Request: bad field: %s", fe.Field), http.StatusBadRequest)
				return
			}
			h.logger.Error("Create user failed", "method", r.Method, "error", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		output := CreateUserResponse{Name: user.Name, Age: user.Age, ID: user.ID}

		data, err := json.Marshal(output)
		if err != nil {
			h.logger.Error("Create user failed: json.Marshal", "method", r.Method, "error", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)

	case http.MethodGet:

		usersCopy, err := h.service.GetAllUsers(ctx)
		if err != nil {
			h.logger.Error("Get all users failed", "method", r.Method, "error", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(usersCopy)
		if err != nil {
			h.logger.Error("Get all users failed: json.Marshal", "method", r.Method, "error", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) UsersByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	switch r.Method {
	case http.MethodGet:
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil || id <= 0 {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(ctx, waitingTime)
		defer cancel()
		outputUser, err := h.service.GetUserByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, context.Canceled):
				h.logger.Info("client went away")
				return
			case errors.Is(err, context.DeadlineExceeded):
				h.logger.Warn("request timed out", "err", err)
				http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
				return
			case errors.Is(err, ErrUserNotFound):
				http.Error(w, "404 Not Found", http.StatusNotFound)
				return
			}
			h.logger.Error("Get user failed", "method", r.Method, "user_id", id, "error", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(outputUser)
		if err != nil {
			h.logger.Error("Get user failed: json.Marshal", "method", r.Method, "user_id", id, "error", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case http.MethodDelete:

		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil || id <= 0 {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		err = h.service.DeleteUser(ctx, id)
		if err != nil {
			if errors.Is(err, ErrUserNotFound) {
				http.Error(w, "404 Not Found", http.StatusNotFound)
				return
			}
			h.logger.Error("Delete user failed", "method", r.Method, "user_id", id, "error", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	case http.MethodPut:

		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil || id <= 0 {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		input := UpdateUserRequest{}
		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		newUser, err := h.service.UpdateUser(ctx, id, input.Name, input.Age)
		if err != nil {
			var fe *FieldError
			if errors.As(err, &fe) {
				http.Error(w, fmt.Sprintf("400 Bad Request: bad field: %s", fe.Field), http.StatusBadRequest)
				return
			}

			if errors.Is(err, ErrUserNotFound) {
				http.Error(w, "404 Not Found", http.StatusNotFound)
				return
			}
			h.logger.Error("Put user failed", "method", r.Method, "user_id", id, "error", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		output := UpdateUserResponse{ID: newUser.ID, Name: newUser.Name, Age: newUser.Age}

		data, err := json.Marshal(output)
		if err != nil {
			h.logger.Error("Put user failed: json.Marshal", "method", r.Method, "user_id", id, "error", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
