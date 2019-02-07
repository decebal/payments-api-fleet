package domain


type BeneficiaryParty struct {
	accountName       string
	accountNumber     string
	accountNumberCode string
	accountType       int
	address           string
	bankId            string
	bankIdCode        string
	name              string
}

type Charge struct {
	amount   float32
	currency string
}

type ChargesInformation struct {
	bearerCode              string
	senderCharges           []Charge
	receiverChargesAmount   float32
	receiverChargesCurrency string
}

type DebtorParty struct {
	accountName       string
	accountNumber     string
	accountNumberCode string
	address           string
	bankId            string
	bankIdCode        string
	name              string
}

type Fx struct {
	contractReference string
	exchangeRate      float32
	originalAmount    float32
	originalCurrency  string
}

type SponsorParty struct {
	accountNumber string
	bankId        string
	bankIdCode    string
}

type Attributes struct {
	amount               int
	beneficiaryParty     BeneficiaryParty
	chargesInformation   ChargesInformation
	currency             string
	debtorParty          DebtorParty
	endToEndReference    string
	fx                   Fx
	numericReference     int
	paymentId            string
	paymentPurpose       string
	paymentScheme        string
	paymentType          string
	processingDate       string
	reference            string
	schemePaymentSubType string
	schemePaymentType    string
	sponsorParty         SponsorParty
}
