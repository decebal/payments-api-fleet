package auth

import (
	"encoding/json"
	"github.com/decebal/payments-api-fleet/api/auth"
	"github.com/decebal/payments-api-fleet/api/errorHandler"
	"github.com/decebal/payments-api-fleet/api/persistence/domain/users"
	"io/ioutil"
	"net/http"
)

type login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginSuccess struct {
	Token string `json:"token" binding:"required"`
}

// DoLogin processes a login request and returns a 200 with the user or a Unauthorised when the login failed
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var u login

	err = json.Unmarshal(body, &u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	found, err := users.GetUserByUsername(u.Username)

	if err != nil {
		errorHandler.OutputHTTPError("username/password incorrect", w, http.StatusUnauthorized)
		return
	}

	token, err := auth.GetJWT(found.Username, found.ID)

	if err != nil {
		errorHandler.OutputHTTPError("username/password incorrect", w, http.StatusUnauthorized)
		return
	}

	t := loginSuccess{Token: token}
	json.NewEncoder(w).Encode(t)
}
