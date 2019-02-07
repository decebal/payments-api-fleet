package domain

import (
	"github.com/mishudark/eventhus"
)

// base Aggregate - Payment
type Payment struct {
	eventhus.BaseAggregate
	OrganisationId int
	Attributes     Attributes
}

// ApplyChange to payment
func (payment *Payment) ApplyChange(event eventhus.Event) {
	switch eventData := event.Data.(type) {
	case *PaymentPerformed:
		payment.Attributes = eventData.Attributes
		payment.OrganisationId = eventData.OrganizationId
		payment.ID = event.AggregateID
	case *PaymentAttributesUpdated:
		payment.Attributes = eventData.Attributes
	case *PaymentRetracted:
		payment.Version = -1 // mark as deleted
	}
}

// HandleCommand create events and validate based on such command
func (payment *Payment) HandleCommand(command eventhus.Command) error {
	event := eventhus.Event{
		AggregateID:   payment.ID,
		AggregateType: "Payment",
	}

	switch commandData := command.(type) {
	case PerformPayment:
		event.AggregateID = commandData.AggregateID
		event.Data = &PaymentPerformed{
			commandData.OrganizationId,
			commandData.Attributes,
		}

	case UpdatePayment:
		event.Data = &PaymentAttributesUpdated{
			commandData.Attributes,
		}

	case DeletePayment:
		event.Data = &PaymentRetracted{}
	}

	payment.BaseAggregate.ApplyChangeHelper(payment, event, true)

	return nil
}
