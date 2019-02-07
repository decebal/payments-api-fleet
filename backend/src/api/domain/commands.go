package domain

import "github.com/mishudark/eventhus"

// PerformPayment to a registered organization
type PerformPayment struct {
	eventhus.BaseCommand
	OrganizationId int
	Attributes Attributes
}

// UpdatePayment, this should be split in more specific use cases, general for now
type UpdatePayment struct {
	eventhus.BaseCommand
	Attributes Attributes
}

// Delete Payment, this should be a corner case,
// e.g. for times where a payment was registered by mistake,
// or unconfirmed by the payment processor
type DeletePayment struct {
	eventhus.BaseCommand
}
