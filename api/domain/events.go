package domain

// AddPayment event
type PaymentPerformed struct {
	OrganizationId string     `json:"organization_id"`
	Attributes     Attributes `json:"attributes"`
}

// UpdatePayment event
type PaymentAttributesUpdated struct {
	Attributes     Attributes `json:"attributes"`
}

// RemovePayment event
type PaymentRetracted struct {}
