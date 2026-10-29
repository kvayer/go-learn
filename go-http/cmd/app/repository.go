package main

import (
	"context"
	"sync"
)

type UserRepository interface {
	Create(ctx context.Context, name string, age int) (User, error)
	GetByID(ctx context.Context, id int) (User, bool, error)
	GetAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, id int, name string, age int) (User, bool, error)
	Delete(ctx context.Context, id int) (bool, error)
}

type MemoryUserRepository struct {
	users  map[int]User
	nextID int
	mu     sync.RWMutex
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:  map[int]User{},
		nextID: 1,
	}
}

func (r *MemoryUserRepository) Create(ctx context.Context, name string, age int) (User, error) {
	if err := ctx.Err(); err != nil {
		return User{}, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	user := User{ID: r.nextID, Name: name, Age: age}
	r.nextID++
	r.users[user.ID] = user
	return user, nil
}

func (r *MemoryUserRepository) GetByID(ctx context.Context, id int) (User, bool, error) {
	if err := ctx.Err(); err != nil {
		return User{}, false, err
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.users[id]
	return user, ok, nil
}

func (r *MemoryUserRepository) GetAll(ctx context.Context) ([]User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	sliceUsers := make([]User, 0, len(r.users))
	for _, user := range r.users {
		sliceUsers = append(sliceUsers, user)
	}
	return sliceUsers, nil
}

func (r *MemoryUserRepository) Update(ctx context.Context, id int, name string, age int) (User, bool, error) {
	if err := ctx.Err(); err != nil {
		return User{}, false, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.users[id]
	if !ok {
		return User{}, false, nil
	}
	newUser := User{
		ID:   id,
		Name: name,
		Age:  age,
	}
	r.users[id] = newUser
	return newUser, true, nil
}

func (r *MemoryUserRepository) Delete(ctx context.Context, id int) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.users[id]
	if !ok {
		return false, nil
	}
	delete(r.users, id)
	return true, nil
}
