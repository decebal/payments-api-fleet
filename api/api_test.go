package main

import (
	"bytes"
	"encoding/json"
	"github.com/decebal/payments-api-fleet/api/auth"
	"github.com/decebal/payments-api-fleet/api/persistence/domain/payments"
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

// Authorization & Authentication
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

// Users
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

// Payments

type NewPaymentRequest struct {
	OrganisationId string              `json:"organisation_id"`
	Attributes     payments.Attributes `json:"attributes"`
}

type PatchPaymentRequest struct {
	ID             string              `json:"id"`
	Attributes     payments.Attributes `json:"attributes"`
}

type PaymentResponse struct {
	ID             string              `json:"id"`
	Created        string              `json:"created"`
	Updated        string              `json:"updated"`
	Type           string              `json:"type"`
	Version        float32             `json:"version"`
	Attributes     payments.Attributes `json:"attributes"`
	OrganisationId string              `json:"organisation_id"`
}

func TestGetPayments(t *testing.T) {
	token, _ := auth.GetJWT("cbrown", 1)

	req, _ := http.NewRequest("GET", "/payments", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		var paymentsList []PaymentResponse
		body, err := ioutil.ReadAll(w.Body)

		if err != nil {
			return false
		}

		err = json.Unmarshal(body, &paymentsList)

		return w.Code == http.StatusOK && err == nil
	})
}

func TestAddPayment(t *testing.T) {
	token, _ := auth.GetJWT("cbrown", 1)

	newPaymentRequest := NewPaymentRequest{
		OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		Attributes: payments.Attributes{
			Amount:   100.21,
			Currency: "GBP",
		},
	}
	jsonData, _ := json.Marshal(newPaymentRequest)

	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		var paymentResponse PaymentResponse

		body, err := ioutil.ReadAll(w.Body)

		if err != nil {
			return false
		}

		err = json.Unmarshal(body, &paymentResponse)

		return w.Code == http.StatusOK &&
			err == nil &&
			newPaymentRequest.Attributes.String() == paymentResponse.Attributes.String()
	})
}

func TestUpdatePaymentById(t *testing.T) {
	token, _ := auth.GetJWT("cbrown", 1)

	paymentUuid := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
	patchPaymentRequest := PatchPaymentRequest{
		ID: paymentUuid,
		Attributes: payments.Attributes{
			Amount:            100.21,
			Currency:          "GBP",
			EndToEndReference: "Wil piano Jan",
			Fx: payments.Fx{
				ContractReference: "FX123",
				ExchangeRate:      2.00000,
				OriginalAmount:    200.42,
				OriginalCurrency:  "USD",
			},
		},
	}
	jsonData, _ := json.Marshal(patchPaymentRequest)

	req, _ := http.NewRequest("PATCH", "/payments/"+paymentUuid, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		var paymentResponse PaymentResponse
		body, err := ioutil.ReadAll(w.Body)

		if err != nil {
			return false
		}

		err = json.Unmarshal(body, &paymentResponse)

		return w.Code == http.StatusOK &&
			err == nil &&
			patchPaymentRequest.Attributes.String() == paymentResponse.Attributes.String()
	})
}

func TestRemovePayment(t *testing.T) {
	token, _ := auth.GetJWT("cbrown", 1)

	paymentUuid := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
	req, _ := http.NewRequest("DELETE", "/payments/"+paymentUuid, nil)

	requiredLen := len(payments.Payments) - 1

	req.Header.Set("Authorization", "Bearer "+token)

	testHTTPResponse(t, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == http.StatusNoContent &&
			len(payments.Payments) == requiredLen
	})
}
