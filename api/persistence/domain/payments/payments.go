package payments

import (
	"errors"
	"github.com/satori/go.uuid"
	"time"
)

var Payments = []Payment{
	Payment{ID:"4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", Version:0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID:"216d4da9-e59a-4cc6-8df3-3da6e7580b77", Version:0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID:"7eb8277a-6c91-45e9-8a03-a27f82aca350", Version:0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID:"97fe60ba-1334-439f-91db-32cc3cde036a", Version:0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID:"ab4bbd28-33c6-4231-9b64-0e96190f59ef", Version:0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID:"7f172f5c-f810-4ebe-b015-cb1fc24c6b66", Version:0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
}

type Payment struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Version  float32 `json:"version"`
	Attributes Attributes `json:"attributes"`
	OrganisationId string `json:"organisation_id"`
	Created  time.Time
	Updated  time.Time
}

type Attributes struct {

}

func getNewPaymentID() string {
	newUuid := uuid.NewV4()

	return newUuid.String()
}

func filter(test func(Payment) bool) (ret []Payment) {
	for _, payment := range Payments {
		if test(payment) {
			ret = append(ret, payment)
		}
	}

	return
}

// GetUserByID returns the user requested by the id, or returns an error
func GetPaymentByID(id string) (Payment, error) {
	var payment Payment

	test := func(u Payment) bool { return u.ID == id }

	p := filter(test)

	if len(p) == 0 {
		return payment, errors.New("payment not found: " + string(id))
	}

	payment = p[0]

	return payment, nil
}

// GetAllPayments from the database and returns them in a list
func GetAllPayments() []Payment {
	return Payments
}

// RemovePayment takes an id and removes that user if found
func RemovePayment(id string) error {
	_, e := GetPaymentByID(id)

	if e != nil {
		return e
	}

	test := func(p Payment) bool { return p.ID == id }
	Payments = filter(test)

	return nil
}

// AddPayment to the main users table
func AddPayment(organisationId string, attributes Attributes) (Payment, error) {
	p := Payment{
		ID: getNewPaymentID(),
		Version: 0,
		Type: "Payment",
		Attributes: attributes,
		Created: time.Now(),
		Updated: time.Now(),
	}

	// any validation over new payments should go in here

	Payments = append(Payments, p)

	return p, nil
}
