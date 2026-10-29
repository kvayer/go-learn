package main

import "context"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name string, age int) (User, error) {
	if name == "" || age <= 0 {
		return User{}, ErrInvalidUserData
	}
	user, err := s.repo.Create(ctx, name, age)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, name string, age int) (User, error) {
	if name == "" || age <= 0 {
		return User{}, ErrInvalidUserData
	}
	user, ok, err := s.repo.Update(ctx, id, name, age)
	if err != nil {
		return user, err
	}
	if !ok {
		return user, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (User, error) {
	user, ok, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return User{}, err
	}
	if !ok {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	ok, err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	if !ok {
		return ErrUserNotFound
	}
	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return users, err
	}
	return users, nil
}
