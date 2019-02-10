package domain


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
	Amount               float32
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
