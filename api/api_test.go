package main

import (
	"bytes"
	"encoding/json"
	"github.com/decebal/payments-api-fleet/api/auth"
	"github.com/decebal/payments-api-fleet/api/persistence/domain/users"
	"github.com/decebal/payments-api-fleet/api/routes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func testHTTPResponse(t *testing.T, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	router := routes.SetupRoutes()
	router.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func TestUnauthorized(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == http.StatusUnauthorized
	})
}

func TestLoginSuccess(t *testing.T) {
	user := LoginRequest{"cbrown", "cbrown123"}
	jsonData, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		var token LoginResponse

		body, err := ioutil.ReadAll(w.Body)

		if err != nil {
			return false
		}

		err = json.Unmarshal(body, &token)

		return w.Code == http.StatusOK && err == nil && len(token.Token) > 0
	})
}

func TestLoginFail(t *testing.T) {
	user := LoginRequest{"Cbrown", "312nworbc"}
	jsonData, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		var errResponse ErrorResponse

		body, err := ioutil.ReadAll(w.Body)

		if err != nil {
			return false
		}

		err = json.Unmarshal(body, &errResponse)

		return w.Code == http.StatusUnauthorized && err == nil && errResponse.Error == "username/password incorrect"
	})
}

func TestGetUsers(t *testing.T) {
	token, _ := auth.GetJWT("cbrown", 1)

	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		var users []UserResponse
		body, err := ioutil.ReadAll(w.Body)

		if err != nil {
			return false
		}

		err = json.Unmarshal(body, &users)

		return w.Code == http.StatusOK && err == nil && users[0].Password == ""
	})
}

func TestAddUser(t *testing.T) {
	newUser := NewUserRequest{"mmouse", "mmouse123"}
	jsonData, _ := json.Marshal(newUser)

	token, _ := auth.GetJWT("cbrown", 1)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		var user UserResponse

		body, err := ioutil.ReadAll(w.Body)

		if err != nil {
			return false
		}

		err = json.Unmarshal(body, &user)

		return w.Code == http.StatusOK && err == nil && newUser.Username == user.Username && user.Password == "" && user.ID == 6
	})
}

func TestAddNewUserWithSameUsername(t *testing.T) {
	newUser := NewUserRequest{"CBROWN", "cbrown123"}
	jsonData, _ := json.Marshal(newUser)

	token, _ := auth.GetJWT("cbrown", 1)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		var errResponse ErrorResponse
		body, err := ioutil.ReadAll(w.Body)

		if err != nil {
			return false
		}

		err = json.Unmarshal(body, &errResponse)

		return w.Code == http.StatusBadRequest && err == nil && errResponse.Error == "username already exists: "+newUser.Username
	})
}

func TestRemoveUser(t *testing.T) {
	token, _ := auth.GetJWT("cbrown", 1)
	req, _ := http.NewRequest("DELETE", "/users/2", nil)

	requiredLen := len(users.Users) - 1

	req.Header.Set("Authorization", "Bearer "+token)

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == http.StatusNoContent && len(users.Users) == requiredLen
	})
}
