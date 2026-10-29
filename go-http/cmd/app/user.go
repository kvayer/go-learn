package main

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateUserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
