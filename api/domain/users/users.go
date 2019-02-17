package users

import (
	"encoding/json"
	"github.com/decebal/payments-api-fleet/api/auth"
	"github.com/decebal/payments-api-fleet/api/errorHandler"
	"github.com/decebal/payments-api-fleet/api/persistence/domain/users"
	"io/ioutil"
	"net/http"
)

type partialUser struct {
	Username string
	Password string
}

// GetUsers returns all the users in the system without the password
func GetUsers(w http.ResponseWriter, r *http.Request, u users.User) {
	if r.Method != "GET" {
		return
	}

	l := users.GetAllUsers()

	json.NewEncoder(w).Encode(l)
}

// AddUser processes a new user request and returns the new user if successful
func AddUser(w http.ResponseWriter, r *http.Request, u users.User) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var p partialUser

	err = json.Unmarshal(body, &p)

	hashPassword := auth.HashAndSalt(p.Password)
	user, err := users.AddUser(p.Username, hashPassword)

	if err != nil {
		errorHandler.OutputHTTPError(err.Error(), w, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// DeleteUser processes the removal of a user and reports 204 on ok
func DeleteUser(w http.ResponseWriter, r *http.Request, u users.User, id int) {
	err := users.RemoveUser(id)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	errorHandler.OutputHTTPError("not a valid user id", w, http.StatusBadRequest)
}
