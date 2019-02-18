package routes

import (
	"github.com/decebal/payments-api-fleet/api/auth"
	domainAuth "github.com/decebal/payments-api-fleet/api/domain/auth"
	"github.com/decebal/payments-api-fleet/api/domain/payments"
	"github.com/decebal/payments-api-fleet/api/domain/users"
	"github.com/decebal/payments-api-fleet/api/errorHandler"
	domainUsers "github.com/decebal/payments-api-fleet/api/persistence/domain/users"
	"net/http"
	"strconv"
	"strings"
)

func shouldAuthorize(fn func(http.ResponseWriter, *http.Request, domainUsers.User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userToken := r.Header.Get("Authorization")

		if userToken == "" {
			errorHandler.OutputHTTPError("not authorized", w, http.StatusUnauthorized)
			return
		}

		userToken = userToken[len("Bearer "):]
		userId, err := auth.DecodeJWT(userToken)
		if err != nil {
			errorHandler.OutputHTTPError("not authorized", w, http.StatusUnauthorized)
			return
		}

		u, err := domainUsers.GetUserByID(userId)
		if err != nil {
			errorHandler.OutputHTTPError("not authorized", w, http.StatusUnauthorized)
			return
		}

		fn(w, r, u)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request, u domainUsers.User) {
	if r.Method == http.MethodPost {
		users.AddUser(w, r, u)
		return
	}

	if r.Method == http.MethodDelete {
		params := strings.Split(r.URL.String(), "/")

		if id, err := strconv.Atoi(params[2]); err == nil {
			users.DeleteUser(w, r, u, id)
			return
		}

		errorHandler.OutputHTTPError("invalid integer", w, http.StatusBadRequest)
		return
	}

	users.GetUsers(w, r, u)
}


func paymentsHandler(w http.ResponseWriter, r *http.Request, u domainUsers.User) {
	if r.Method == http.MethodPost {
		payments.AddPayment(w, r, u)
		return
	}

	if r.Method == http.MethodPatch {
		params := strings.Split(r.URL.String(), "/")

		if len(params) == 3 {
			payments.UpdatePayment(w, r, u, params[2])
			return
		}

		errorHandler.OutputHTTPError("Payment ID not provided", w, http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodDelete {
		params := strings.Split(r.URL.String(), "/")

		if len(params) == 3 {
			payments.DeletePayment(w, r, u, params[2])
			return
		}

		errorHandler.OutputHTTPError("Payment ID not provided", w, http.StatusBadRequest)
		return
	}

	payments.GetPayments(w, r, u)
}

// SetupRoutes in gin
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/login", domainAuth.HandleLogin)
	mux.HandleFunc("/users", shouldAuthorize(userHandler))
	mux.HandleFunc("/users/", shouldAuthorize(userHandler))
	mux.HandleFunc("/payments", shouldAuthorize(paymentsHandler))
	mux.HandleFunc("/payments/", shouldAuthorize(paymentsHandler))

	return mux
}
