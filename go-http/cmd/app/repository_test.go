package main

import (
	"context"
	"errors"
	"testing"
)

func TestMemoryUserRepository_Create_CancelledContext(t *testing.T) {
	repo := NewMemoryUserRepository()

	// проверка создания при отмененном контексте
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // отмена контекста

	user, err := repo.Create(ctx, "mike", 22) // вызов создания

	expectedUser := User{}

	if user != expectedUser { // если польватель не пустой, тогда выкидываем ошибку, тк пользователь был создан
		t.Fatalf("expected %+v, got %+v", expectedUser, user)
	}

	if !errors.Is(err, context.Canceled) { // если итог ошибки вызова Create было не отменой контекста, тогда выбрасываем ошибку
		t.Fatalf("expected %v, got %v", context.Canceled, err)
	}
	// создание нового контекста для проверки, были ли созданы пользователи
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	users, err := repo.GetAll(ctx) // получаем всех пользователей

	if err != nil { // если произошла ошибка при получении пользователей при новом контексте - выбрасываем ошибку
		t.Fatalf("expected: nil, got: %v", err)
	}

	if len(users) != 0 { // если пользователь был создан - выкидываем ошибку
		t.Fatalf("expected: 0, got: %d", len(users))
	}

	user, err = repo.Create(ctx, "mike", 22)

	if err != nil { // если произошла ошибка при создании - выбрасываем ошибку
		t.Fatalf("expected: nil, got: %v", err)
	}

	if user.ID != 1 { // новый пользователь должен иметь id 1, поскольку до этого пользователей не было создано
		t.Fatalf("expected: 1, got: %d", user.ID)
	}
}

func TestMemoryUserRepository_GetByID(t *testing.T) {

	repo := NewMemoryUserRepository()
	activeCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cancelledCtx, cancel2 := context.WithCancel(context.Background())
	cancel2()
	expectedUser, err := repo.Create(activeCtx, "Mike", 22)
	if err != nil {
		t.Fatalf("couldn't Create user: %v", err)
	}

	testCases := []struct {
		name         string
		ctx          context.Context
		id           int
		expectedUser User
		expectedOk   bool
		expectedErr  error
	}{
		{name: "found", ctx: activeCtx, id: expectedUser.ID, expectedUser: expectedUser, expectedOk: true, expectedErr: nil},
		{name: "not found", ctx: activeCtx, id: 999, expectedUser: User{}, expectedOk: false, expectedErr: nil},
		{name: "cancelled context", ctx: cancelledCtx, id: expectedUser.ID, expectedUser: User{}, expectedOk: false, expectedErr: context.Canceled},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			user, ok, err := repo.GetByID(test.ctx, test.id)
			if user != test.expectedUser {
				t.Fatalf("expected user: %+v, got: %+v", test.expectedUser, user)
			}
			if ok != test.expectedOk {
				t.Fatalf("expected ok: %t, got: %t", test.expectedOk, ok)
			}
			if !errors.Is(err, test.expectedErr) {
				t.Fatalf("expected err: %v, got: %v", test.expectedErr, err)
			}
		})
	}
}

func TestMemoryUserRepository_Update_CancelledContext(t *testing.T) {
	repo := NewMemoryUserRepository()
	ctx := context.Background()
	createdUser, err := repo.Create(ctx, "Mike", 22)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	ctxCanceled, cancel := context.WithCancel(context.Background())
	cancel()
	returnedUser, ok, err := repo.Update(ctxCanceled, createdUser.ID, "Anna", 24)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected: context canceled, got: %v", err)
	}
	if ok {
		t.Fatal("expedted: false, got: true")
	}
	if returnedUser != (User{}) {
		t.Fatalf("expected zero User, got: %+v", returnedUser)
	}
	storedUser, ok, err := repo.GetByID(ctx, createdUser.ID)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	if !ok {
		t.Fatalf("expedted: true, got: %t", ok)
	}
	if createdUser != storedUser {
		t.Fatalf("expected User: %+v, got: %+v", createdUser, storedUser)
	}
}

func TestMemoryUserRepository_Update_NotFound(t *testing.T) {
	repo := NewMemoryUserRepository() //новый repository
	ctx := context.Background()       //действующий context

	createdUser, err := repo.Create(ctx, "Mike", 22) //созданный Mike
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	gotUser, ok, err := repo.Update(ctx, 999, "Anna", 24) //попытка обновить ID 999
	//проверка User{}, false, nil
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if ok {
		t.Fatal("expected: false, got: true")
	}

	if gotUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", gotUser)
	}

	storedUser, ok, err := repo.GetByID(ctx, createdUser.ID)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !ok {
		t.Fatal("expected: true, got: false")
	}

	if storedUser != createdUser { //исходный Mike остался неизменным
		t.Fatalf("expected: %+v, got: %+v", createdUser, storedUser)
	}

	missingUser, ok, err := repo.GetByID(ctx, 999)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if ok {
		t.Fatal("expected: false, got: true")
	}

	if missingUser != (User{}) {
		t.Fatalf("expected: User{}, got: %+v", missingUser)
	}
}

func TestMemoryUserRepository_Update_Success(t *testing.T) {

	repo := NewMemoryUserRepository()
	ctx := context.Background()
	createdUser, err := repo.Create(ctx, "Mike", 22)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	changeUser, ok, err := repo.Update(ctx, createdUser.ID, "Anna", 24)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !ok {
		t.Fatal("expected: true, got: false")
	}
	expectedUser := User{ID: createdUser.ID, Name: "Anna", Age: 24}
	if changeUser != expectedUser {
		t.Fatalf("expected: %v, got: %+v", expectedUser, changeUser)
	}

	storedUser, ok, err := repo.GetByID(ctx, createdUser.ID)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !ok {
		t.Fatal("expected: true, got: false")
	}

	if storedUser != changeUser {
		t.Fatalf("expected: %+v, got: %+v", changeUser, storedUser)
	}
}

func TestMemoryUserRepository_Delete_CancelledContext(t *testing.T) {
	repo := NewMemoryUserRepository()
	contextCanceled, cancel := context.WithCancel(context.Background())
	cancel()
	ctx := context.Background()
	createdUser1, err := repo.Create(ctx, "Mike", 22)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	createdUser2, err := repo.Create(ctx, "Anna", 24)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	ok, err := repo.Delete(contextCanceled, createdUser1.ID)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected: %v, got: %v", context.Canceled, err)
	}

	if ok {
		t.Fatal("expected: false, got: true")
	}

	storedUser1, ok, err := repo.GetByID(ctx, createdUser1.ID)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !ok {
		t.Fatal("expected: true, got: false")
	}

	if storedUser1 != createdUser1 {
		t.Fatalf("expected: %+v, got: %+v", createdUser1, storedUser1)
	}

	storedUser2, ok, err := repo.GetByID(ctx, createdUser2.ID)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !ok {
		t.Fatal("expected: true, got: false")
	}

	if storedUser2 != createdUser2 {
		t.Fatalf("expected: %+v, got: %+v", createdUser2, storedUser2)
	}
}

func TestMemoryUserRepository_Delete_NotFound(t *testing.T) {
	repo := NewMemoryUserRepository()
	ctx := context.Background()
	createdUser1, err := repo.Create(ctx, "Mike", 22)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	createdUser2, err := repo.Create(ctx, "Anna", 24)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	ok, err := repo.Delete(ctx, 999)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if ok {
		t.Fatal("expected: false, got: true")
	}

	storedUser1, ok, err := repo.GetByID(ctx, createdUser1.ID)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !ok {
		t.Fatal("expected: true, got: false")
	}

	if storedUser1 != createdUser1 {
		t.Fatalf("expected: %+v, got: %+v", createdUser1, storedUser1)
	}

	storedUser2, ok, err := repo.GetByID(ctx, createdUser2.ID)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !ok {
		t.Fatal("expected: true, got: false")
	}

	if storedUser2 != createdUser2 {
		t.Fatalf("expected: %+v, got: %+v", createdUser2, storedUser2)
	}
}

func TestMemoryUserRepository_Delete_Success(t *testing.T) {
	repo := NewMemoryUserRepository()
	ctx := context.Background()
	createdUser1, err := repo.Create(ctx, "Mike", 22)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	createdUser2, err := repo.Create(ctx, "Anna", 24)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	ok, err := repo.Delete(ctx, createdUser1.ID)

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !ok {
		t.Fatal("expected: true, got: false")
	}

	expectedUser1 := User{}
	storedUser1, ok, err := repo.GetByID(ctx, createdUser1.ID)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if ok {
		t.Fatal("expected: false, got: true")
	}

	if storedUser1 != expectedUser1 {
		t.Fatalf("expected: %+v, got: %+v", expectedUser1, storedUser1)
	}

	storedUser2, ok, err := repo.GetByID(ctx, createdUser2.ID)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if !ok {
		t.Fatal("expected: true, got: false")
	}

	if storedUser2 != createdUser2 {
		t.Fatalf("expected: %+v, got: %+v", createdUser2, storedUser2)
	}
}
