package main

import (
	"context"
	"fmt"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name string, age int) (User, error) {

	// валидация name и age
	fe := &FieldError{}
	if name == "" {
		fe.Field = "name"
		return User{}, fmt.Errorf("%w: %w", ErrInvalidUserData, fe)
	}
	if age <= 0 {
		fe.Field = "age"
		return User{}, fmt.Errorf("%w: %w", ErrInvalidUserData, fe)
	}

	user, err := s.repo.Create(ctx, name, age)
	if err != nil {
		return user, fmt.Errorf("CreateUser: %w", err)
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, name string, age int) (User, error) {

	if name == "" {
		return User{}, fmt.Errorf("%w: %w", ErrInvalidUserData, &FieldError{Field: "name"})
	}
	if age <= 0 {
		return User{}, fmt.Errorf("%w: %w", ErrInvalidUserData, &FieldError{Field: "age"})
	}

	user, ok, err := s.repo.Update(ctx, id, name, age)
	if err != nil {
		return user, fmt.Errorf("UpdateUser: %w", err)
	}
	if !ok {
		return user, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (User, error) {
	user, ok, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return User{}, fmt.Errorf("GetUserByID: %w", err)
	}
	if !ok {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	ok, err := s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("DeleteUser: %w", err)
	}
	if !ok {
		return ErrUserNotFound
	}
	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return users, fmt.Errorf("GetAllUsers: %w", err)
	}
	return users, nil
}
