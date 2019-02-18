package payments

import (
	"encoding/json"
	"github.com/decebal/payments-api-fleet/api/errorHandler"
	"github.com/decebal/payments-api-fleet/api/persistence/domain/payments"
	"github.com/decebal/payments-api-fleet/api/persistence/domain/users"
	"io/ioutil"
	"net/http"
)

type partialPayment struct {
	ID             string              `json:"id"`
	OrganisationId string              `json:"organisation_id"`
	Attributes     payments.Attributes `json:"attributes"`
}

// GetPayments returns all the payments, eventually this should done be based on organisationId
func GetPayments(w http.ResponseWriter, r *http.Request, u users.User) {
	if r.Method != "GET" {
		return
	}

	l := payments.GetAllPayments()

	err := json.NewEncoder(w).Encode(l)
	if err != nil {
		errorHandler.OutputHTTPError(err.Error(), w, http.StatusInternalServerError)
		return
	}
}

// AddPayment processes a new payment request and returns the new payment if successful
func AddPayment(w http.ResponseWriter, r *http.Request, u users.User) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var p partialPayment

	err = json.Unmarshal(body, &p)

	newPayment, err := payments.AddPayment(p.OrganisationId, p.Attributes)

	if err != nil {
		errorHandler.OutputHTTPError(err.Error(), w, http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(newPayment)

	if err != nil {
		errorHandler.OutputHTTPError(err.Error(), w, http.StatusInternalServerError)
		return
	}
}

// UpdatePayment processes a payment details request and returns the payment if successful
func UpdatePayment(w http.ResponseWriter, r *http.Request, u users.User, id string) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var p partialPayment
	err = json.Unmarshal(body, &p)

	updatePayment, err := payments.UpdatePayment(id, p.Attributes)

	if err != nil {
		errorHandler.OutputHTTPError(err.Error(), w, http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(updatePayment)
	if err != nil {
		errorHandler.OutputHTTPError(err.Error(), w, http.StatusInternalServerError)
		return
	}
}

// DeletePayment processes the removal of a payment and reports 204 on ok
func DeletePayment(w http.ResponseWriter, r *http.Request, u users.User, id string) {
	err := payments.RemovePayment(id)

	if err == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	errorHandler.OutputHTTPError("not a valid Payment Id", w, http.StatusBadRequest)
}
