package main

import (
	"context"
	"errors"
	"slices"
	"testing"
)

var errRepository = errors.New("repository error")

type stubUserRepository struct {
	getAllUsers  []User
	getAllErr    error
	updateCalled bool
	updateUser   User
	updateOK     bool
	updateErr    error
	getByIDUser  User
	getByIDOK    bool
	getByIDErr   error
	createCalled bool
	createUser   User
	createErr    error
	deleteOK     bool
	deleteErr    error
}

func (s *stubUserRepository) Delete(ctx context.Context, id int) (bool, error) {
	return s.deleteOK, s.deleteErr
}

func (s *stubUserRepository) Create(ctx context.Context, name string, age int) (User, error) {
	s.createCalled = true
	return s.createUser, s.createErr
}

func (s *stubUserRepository) GetByID(ctx context.Context, id int) (User, bool, error) {
	return s.getByIDUser, s.getByIDOK, s.getByIDErr
}

func (s *stubUserRepository) Update(ctx context.Context, id int, name string, age int) (User, bool, error) {
	s.updateCalled = true
	return s.updateUser, s.updateOK, s.updateErr
}

func (s *stubUserRepository) GetAll(ctx context.Context) ([]User, error) {
	return s.getAllUsers, s.getAllErr
}

func TestUserService_DeleteUser_Success(t *testing.T) {
	repo := stubUserRepository{deleteOK: true}
	ctx := context.Background()

	service := NewUserService(&repo)
	err := service.DeleteUser(ctx, 1)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
}

func TestUserService_DeleteUser_NotFound(t *testing.T) {
	repo := stubUserRepository{deleteOK: false, deleteErr: nil}
	ctx := context.Background()

	service := NewUserService(&repo)
	err := service.DeleteUser(ctx, 1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("expected: %v, got: %v", ErrUserNotFound, err)
	}
}

func TestUserService_DeleteUser_RepositoryError(t *testing.T) {
	repo := stubUserRepository{deleteOK: false, deleteErr: errRepository}
	ctx := context.Background()

	service := NewUserService(&repo)
	err := service.DeleteUser(ctx, 1)
	if !errors.Is(err, errRepository) {
		t.Fatalf("expected: %v, got: %v", errRepository, err)
	}
}

func TestUserService_CreateUser_EmptyName(t *testing.T) {
	repo := stubUserRepository{}
	ctx := context.Background()

	service := NewUserService(&repo)
	createdUser, err := service.CreateUser(ctx, "", 24)
	if repo.createCalled {
		t.Fatalf("repository Create shouldn't be called")
	}
	if !errors.Is(err, ErrInvalidUserData) {
		t.Fatalf("expected: %v, got: %v", ErrInvalidUserData, err)
	}
	if createdUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", createdUser)
	}
}

func TestUserService_CreateUser_InvalidAge(t *testing.T) {
	repo := stubUserRepository{}
	ctx := context.Background()

	service := NewUserService(&repo)
	createdUser, err := service.CreateUser(ctx, "Mike", 0)
	if repo.createCalled {
		t.Fatalf("repository Create shouldn't be called")
	}
	if !errors.Is(err, ErrInvalidUserData) {
		t.Fatalf("expected: %v, got: %v", ErrInvalidUserData, err)
	}
	if createdUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", createdUser)
	}
}

func TestUserService_CreateUser_RepositoryError(t *testing.T) {
	repo := stubUserRepository{createErr: errRepository}
	ctx := context.Background()

	service := NewUserService(&repo)
	createdUser, err := service.CreateUser(ctx, "Mike", 22)
	if !errors.Is(err, errRepository) {
		t.Fatalf("expected: %v, got: %v", errRepository, err)
	}
	if createdUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", createdUser)
	}
}

func TestUserService_CreateUser_Success(t *testing.T) {
	expectedUser := User{ID: 42, Name: "Mike", Age: 22}
	repo := stubUserRepository{createUser: expectedUser}
	ctx := context.Background()

	service := NewUserService(&repo)
	createdUser, err := service.CreateUser(ctx, "Mike", 22)
	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	if createdUser != expectedUser {
		t.Fatalf("expected: %+v got: %+v", expectedUser, createdUser)
	}
}

func TestUserService_GetUserByID_NotFound(t *testing.T) {
	repo := stubUserRepository{}
	ctx := context.Background()

	service := NewUserService(&repo)
	storedUser, err := service.GetUserByID(ctx, 42)
	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("expected: %v, got: %v", ErrUserNotFound, err)
	}
	if storedUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", storedUser)
	}
}

func TestUserService_GetUserByID_RepositoryError(t *testing.T) {
	repo := stubUserRepository{getByIDErr: errRepository}
	ctx := context.Background()

	service := NewUserService(&repo)
	storedUser, err := service.GetUserByID(ctx, 42)
	if !errors.Is(err, errRepository) {
		t.Fatalf("expected: %v, got: %v", errRepository, err)
	}
	if storedUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", storedUser)
	}
}

func TestUserService_GetUserByID_Success(t *testing.T) {
	expectedUser := User{ID: 1, Name: "Mike", Age: 22}
	repo := stubUserRepository{getByIDUser: expectedUser, getByIDOK: true}
	ctx := context.Background()

	service := NewUserService(&repo)
	storedUser, err := service.GetUserByID(ctx, 1)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	if storedUser != expectedUser {
		t.Fatalf("expected: %+v, got: %+v", expectedUser, storedUser)
	}
}

func TestUserService_UpdateUser_EmptyName(t *testing.T) {
	repo := stubUserRepository{}
	ctx := context.Background()

	service := NewUserService(&repo)

	gotUser, err := service.UpdateUser(ctx, 42, "", 22)
	if repo.updateCalled {
		t.Fatal("repository Update shouldn't be called")
	}
	if !errors.Is(err, ErrInvalidUserData) {
		t.Fatalf("expected: %v, got: %v", ErrInvalidUserData, err)
	}

	if gotUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", gotUser)
	}
}

func TestUserService_UpdateUser_InvalidAge(t *testing.T) {
	repo := stubUserRepository{}
	ctx := context.Background()

	service := NewUserService(&repo)

	gotUser, err := service.UpdateUser(ctx, 42, "Mike", 0)
	if repo.updateCalled {
		t.Fatal("repository Update shouldn't be called")
	}
	if !errors.Is(err, ErrInvalidUserData) {
		t.Fatalf("expected: %v, got: %v", ErrInvalidUserData, err)
	}

	if gotUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", gotUser)
	}
}

func TestUserService_UpdateUser_NotFound(t *testing.T) {
	repo := stubUserRepository{}
	ctx := context.Background()

	service := NewUserService(&repo)

	gotUser, err := service.UpdateUser(ctx, 42, "Mike", 22)
	if !repo.updateCalled {
		t.Fatal("repository Update should be called")
	}
	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("expected: %v, got: %v", ErrUserNotFound, err)
	}

	if gotUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", gotUser)
	}
}

func TestUserService_UpdateUser_RepositoryError(t *testing.T) {
	repo := stubUserRepository{updateErr: errRepository}
	ctx := context.Background()

	service := NewUserService(&repo)

	gotUser, err := service.UpdateUser(ctx, 42, "Mike", 22)
	if !repo.updateCalled {
		t.Fatal("repository Update should be called")
	}
	if !errors.Is(err, errRepository) {
		t.Fatalf("expected: %v, got: %v", errRepository, err)
	}

	if gotUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", gotUser)
	}
}

func TestUserService_UpdateUser_Success(t *testing.T) {
	expectedUser := User{ID: 42, Name: "Mike", Age: 22}
	repo := stubUserRepository{updateUser: expectedUser, updateOK: true}
	ctx := context.Background()

	service := NewUserService(&repo)

	gotUser, err := service.UpdateUser(ctx, 42, "Mike", 22)
	if !repo.updateCalled {
		t.Fatal("repository Update should be called")
	}
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if gotUser != expectedUser {
		t.Fatalf("expected: %+v, got: %+v", expectedUser, gotUser)
	}
}

func TestUserService_GetAllUsers_RepositoryError(t *testing.T) {
	repo := stubUserRepository{getAllErr: errRepository}
	ctx := context.Background()

	service := NewUserService(&repo)

	users, err := service.GetAllUsers(ctx)
	if !errors.Is(err, errRepository) {
		t.Fatalf("expected: %v, got: %v", errRepository, err)
	}

	expectedUsers := make([]User, 0)
	if !slices.Equal(users, expectedUsers) {
		t.Fatalf("expected: %+v, got: %+v", expectedUsers, users)
	}
}

func TestUserService_GetAllUsers_Success(t *testing.T) {
	expectedUsers := []User{{ID: 42, Name: "Mike", Age: 22}}
	repo := stubUserRepository{getAllUsers: expectedUsers}
	ctx := context.Background()

	service := NewUserService(&repo)

	users, err := service.GetAllUsers(ctx)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !slices.Equal(users, expectedUsers) {
		t.Fatalf("expected: %+v, got: %+v", expectedUsers, users)
	}
}
