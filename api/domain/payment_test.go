// +build unit

package domain

import (
	"github.com/mishudark/eventhus"
	. "github.com/onsi/ginkgo"
)


var _ = Describe("RegisterPayment", func() {
	reg := eventhus.NewEventRegister()
	reg.Set(PaymentPerformed{})

})

var _ = Describe("UpdatePayment", func() {

})

var _ = Describe("RemovePayment", func() {

})

var _ = Describe("ListPayments", func() {

})
