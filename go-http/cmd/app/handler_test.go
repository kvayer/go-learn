package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
)

func TestUserHandler_Users_POST_Success(t *testing.T) {
	body := strings.NewReader(`{"name":"Mike","age":22}`)
	request := httptest.NewRequest(http.MethodPost, "/users", body)
	recorder := httptest.NewRecorder()

	expectedUser := User{ID: 42, Name: "Mike", Age: 22}
	repo := stubUserRepository{createUser: expectedUser}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)
	handler.Users(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status: %d, got: %d", http.StatusCreated, recorder.Code)
	}
	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("expected: application/json, got: %s", contentType)
	}

	expectedResponse := CreateUserResponse{
		ID:   expectedUser.ID,
		Name: expectedUser.Name,
		Age:  expectedUser.Age,
	}

	var gotResponse CreateUserResponse
	err := json.NewDecoder(recorder.Body).Decode(&gotResponse)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if gotResponse != expectedResponse {
		t.Fatalf("expected: %+v, got:%+v", expectedResponse, gotResponse)
	}
}

func TestUserHandler_Users_POST_InvalidJSON(t *testing.T) {
	body := strings.NewReader(`{"name": ;`)
	request := httptest.NewRequest(http.MethodPost, "/users", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)
	handler.Users(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"

	if gotBody != expectedBody {
		t.Fatalf("expected body: %q, got: %q", expectedBody, gotBody)
	}
}

func TestUserHandler_Users_POST_InvalidUserData(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike","age": 0}`)
	request := httptest.NewRequest(http.MethodPost, "/users", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)
	handler.Users(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}

	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"

	if gotBody != expectedBody {
		t.Fatalf("expected body: %s, got: %s", expectedBody, gotBody)
	}

}

func TestUserHandler_Users_POST_RepositoryError(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike","age": 22}`)
	request := httptest.NewRequest(http.MethodPost, "/users", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{
		createErr: errRepository,
	}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)
	handler.Users(recorder, request)
	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected: %d, got: %d", http.StatusInternalServerError, recorder.Code)
	}

	gotBody := recorder.Body.String()
	expectedBody := "500 Internal Server Error\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %s, got: %s", expectedBody, gotBody)
	}
}

func TestUserHandler_Users_GET_Success(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users", nil)
	recorder := httptest.NewRecorder()

	expectedUsers := []User{{ID: 42, Name: "Mike", Age: 22}, {ID: 98, Name: "Anna", Age: 24}}

	repo := stubUserRepository{getAllUsers: expectedUsers}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)
	handler.Users(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected: %d, got: %d", http.StatusOK, recorder.Code)
	}
	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("expected: %s, got: %s", "application/json", contentType)
	}
	var gotUsers []User
	err := json.NewDecoder(recorder.Body).Decode(&gotUsers)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	if !slices.Equal(gotUsers, expectedUsers) {
		t.Fatalf("expected: %+v, got: %+v", expectedUsers, gotUsers)
	}
}

func TestUserHandler_Users_GET_RepositoryError(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{getAllErr: errRepository}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)
	handler.Users(recorder, request)
	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected: %d, got: %d", http.StatusInternalServerError, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "500 Internal Server Error\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %s, got: %s", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_GET_RepositoryError(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users/42", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{getByIDErr: errRepository}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "42")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected: %d, got: %d", http.StatusInternalServerError, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "500 Internal Server Error\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_GET_NotFound(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users/42", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "42")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected: %d, got: %d", http.StatusNotFound, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "404 Not Found\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_GET_Success(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users/44", nil)
	recorder := httptest.NewRecorder()

	expectedUser := User{ID: 44, Name: "Mike", Age: 22}
	repo := stubUserRepository{getByIDOK: true, getByIDUser: expectedUser}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected: %d, got: %d", http.StatusOK, recorder.Code)
	}
	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("expected: %s, got %s", "application/json", contentType)
	}
	var gotUser User
	err := json.NewDecoder(recorder.Body).Decode(&gotUser)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	if gotUser != expectedUser {
		t.Fatalf("expected: %v, got: %v", expectedUser, gotUser)
	}
}

func TestUserHandler_UsersByID_GET_InvalidID_Zero(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users/0", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "0")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_GET_InvalidID_NotNumber(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users/rgf", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "rgf")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_DELETE_InvalidID_Zero(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/users/0", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "0")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_DELETE_InvalidID_NotNumber(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/users/sdfg", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "sdfg")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_DELETE_NotFound(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/users/44", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected: %d, got: %d", http.StatusNotFound, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "404 Not Found\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_DELETE_RepositoryError(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/users/44", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{deleteErr: errRepository}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected: %d, got: %d", http.StatusInternalServerError, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "500 Internal Server Error\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_DELETE_Success(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/users/44", nil)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{deleteOK: true}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected: %d, got: %d", http.StatusNoContent, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := ""
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_RepositoryError(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike", "age": 22}`)
	request := httptest.NewRequest(http.MethodPut, "/users/44", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{updateErr: errRepository}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected: %d, got: %d", http.StatusInternalServerError, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "500 Internal Server Error\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_InvalidJSON(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike", "age": 22`)
	request := httptest.NewRequest(http.MethodPut, "/users/44", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_InvalidUserData_EmptyObject(t *testing.T) {
	body := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodPut, "/users/44", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_InvalidUserData_MissingAge(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike"}`)
	request := httptest.NewRequest(http.MethodPut, "/users/44", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_InvalidID_Zero(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike", "age": 22}`)
	request := httptest.NewRequest(http.MethodPut, "/users/0", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "0")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_InvalidID_NotNumber(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike", "age": 22}`)
	request := httptest.NewRequest(http.MethodPut, "/users/fdsg", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "fdsg")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_NotFound(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike", "age": 22}`)
	request := httptest.NewRequest(http.MethodPut, "/users/44", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected: %d, got: %d", http.StatusNotFound, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "404 Not Found\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_InvalidName(t *testing.T) {
	body := strings.NewReader(`{"name": "", "age": 22}`)
	request := httptest.NewRequest(http.MethodPut, "/users/44", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_InvalidAge(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike", "age": 0}`)
	request := httptest.NewRequest(http.MethodPut, "/users/44", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_InvalidNameAndAge(t *testing.T) {
	body := strings.NewReader(`{"name": "", "age": 0}`)
	request := httptest.NewRequest(http.MethodPut, "/users/44", body)
	recorder := httptest.NewRecorder()

	repo := stubUserRepository{}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected: %d, got: %d", http.StatusBadRequest, recorder.Code)
	}
	gotBody := recorder.Body.String()
	expectedBody := "400 Bad Request\n"
	if gotBody != expectedBody {
		t.Fatalf("expected: %v, got: %v", expectedBody, gotBody)
	}
}

func TestUserHandler_UsersByID_PUT_Success(t *testing.T) {
	body := strings.NewReader(`{"name": "Mike", "age": 22}`)
	request := httptest.NewRequest(http.MethodPut, "/users/44", body)
	recorder := httptest.NewRecorder()

	expectedUser := User{ID: 44, Name: "Mike", Age: 22}
	repo := stubUserRepository{updateUser: expectedUser, updateOK: true}
	service := NewUserService(&repo)
	handler := NewUserHandler(service)

	request.SetPathValue("id", "44")
	handler.UsersByID(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected: %d, got: %d", http.StatusOK, recorder.Code)
	}
	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("expected: %v, got: %v", "application/json", contentType)
	}
	var gotUser User
	err := json.NewDecoder(recorder.Body).Decode(&gotUser)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	if gotUser != expectedUser {
		t.Fatalf("expected: %v, got: %v", expectedUser, gotUser)
	}
}

func TestRouter_UsersByID_GET_Success(t *testing.T) {
	body := strings.NewReader(`{"name":"Mike","age":22}`)
	request := httptest.NewRequest(http.MethodPost, "/users", body)
	recorder := httptest.NewRecorder()

	repo := NewMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	mux := NewRouter(handler)
	mux.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected: %d, got: %d", http.StatusCreated, recorder.Code)
	}
	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("expected: %s, got: %s", "application/json", contentType)
	}
	var expectedUser User
	err := json.NewDecoder(recorder.Body).Decode(&expectedUser)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	request = httptest.NewRequest(http.MethodGet, "/users/1", nil)
	recorder = httptest.NewRecorder()
	mux.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected: %d, got: %d", http.StatusOK, recorder.Code)
	}
	contentType = recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("expected: %s, got: %s", "application/json", contentType)
	}
	var gotUser User
	err = json.NewDecoder(recorder.Body).Decode(&gotUser)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	if gotUser != expectedUser {
		t.Fatalf("expected: %+v, got: %+v", expectedUser, gotUser)
	}
}

func TestHTTP_Users_POST_Success(t *testing.T) {
	repo := NewMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)
	mux := NewRouter(handler)

	server := httptest.NewServer(mux)
	defer server.Close()

	body := strings.NewReader(`{"name":"Mike","age":22}`)
	response, err := server.Client().Post(server.URL+"/users", "application/json", body)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("expected: %d, got: %d", http.StatusCreated, response.StatusCode)
	}
	contentType := response.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("expected: %s, got: %s", "application/json", contentType)
	}
	var gotUser User
	err = json.NewDecoder(response.Body).Decode(&gotUser)
	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	expectedUser := User{ID: 1, Name: "Mike", Age: 22}
	if gotUser != expectedUser {
		t.Fatalf("expected: %+v, got: %+v", expectedUser, gotUser)
	}
}
