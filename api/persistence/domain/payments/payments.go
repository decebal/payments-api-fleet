package payments

import (
	"encoding/json"
	"errors"
	"github.com/imdario/mergo"
	"github.com/satori/go.uuid"
	"time"
)

var Payments = []Payment{
	Payment{ID: "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", Version: 0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID: "216d4da9-e59a-4cc6-8df3-3da6e7580b77", Version: 0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID: "7eb8277a-6c91-45e9-8a03-a27f82aca350", Version: 0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID: "97fe60ba-1334-439f-91db-32cc3cde036a", Version: 0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID: "ab4bbd28-33c6-4231-9b64-0e96190f59ef", Version: 0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
	Payment{ID: "7f172f5c-f810-4ebe-b015-cb1fc24c6b66", Version: 0, OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",},
}

type BeneficiaryParty struct {
	AccountName       string
	AccountNumber     string
	AccountNumberCode string
	AccountType       int
	Address           string
	BankId            string
	BankIdCode        string
	Name              string
}

type Charge struct {
	Amount   float32
	Currency string
}

type ChargesInformation struct {
	BearerCode              string
	SenderCharges           []Charge
	ReceiverChargesAmount   float32
	ReceiverChargesCurrency string
}

type DebtorParty struct {
	AccountName       string
	AccountNumber     string
	AccountNumberCode string
	Address           string
	BankId            string
	BankIdCode        string
	Name              string
}

type Fx struct {
	ContractReference string
	ExchangeRate      float32
	OriginalAmount    float32
	OriginalCurrency  string
}

type SponsorParty struct {
	AccountNumber string
	BankId        string
	BankIdCode    string
}

type Attributes struct {
	Amount               float64
	BeneficiaryParty     BeneficiaryParty
	ChargesInformation   ChargesInformation
	Currency             string
	DebtorParty          DebtorParty
	EndToEndReference    string
	Fx                   Fx
	NumericReference     int
	PaymentId            string
	PaymentPurpose       string
	PaymentScheme        string
	PaymentType          string
	ProcessingDate       string
	Reference            string
	SchemePaymentSubType string
	SchemePaymentType    string
	SponsorParty         SponsorParty
}

func (a *Attributes) String() string {
	jsonData, _ := json.Marshal(a)
	return string(jsonData)
}

type Payment struct {
	ID             string
	Type           string
	Version        float32
	Attributes     Attributes
	OrganisationId string
	Created        time.Time
	Updated        time.Time
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

	test := func(p Payment) bool { return p.ID != id }
	Payments = filter(test)

	return nil
}

// AddPayment to the main payments list
func AddPayment(organisationId string, attributes Attributes) (Payment, error) {
	p := Payment{
		ID:             getNewPaymentID(),
		Version:        0,
		Type:           "Payment",
		Attributes:     attributes,
		OrganisationId: organisationId,
		Created:        time.Now(),
		Updated:        time.Now(),
	}

	// any validation over new payments should go in here

	Payments = append(Payments, p)

	return p, nil
}

// UpdatePayment from the existing list
func UpdatePayment(id string, attributes Attributes) (Payment, error) {
	existingPayment, err := GetPaymentByID(id)

	if err != nil {
		return Payment{}, err
	}

	err = mergo.Merge(&attributes, existingPayment.Attributes)

	if err != nil {
		return Payment{}, err
	}

	p := Payment{
		Version:        existingPayment.Version + 1,
		Attributes:     attributes,
		Updated:        time.Now(),
		OrganisationId: existingPayment.OrganisationId,
		Type:           existingPayment.Type,
		ID:             existingPayment.ID,
		Created:        existingPayment.Created,
	}

	// any validation over payment details should go in here

	var auxPayments []Payment

	test := func(p Payment) bool { return p.ID == id }

	for _, payment := range Payments {
		if test(payment) {
			auxPayments = append(auxPayments, p)
		} else {
			auxPayments = append(auxPayments, payment)
		}
	}
	Payments = auxPayments

	return p, nil
}
