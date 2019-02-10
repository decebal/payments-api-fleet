package main

import (
	"flag"
	"github.com/decebal/payments-api-fleet/api/domain"
	"github.com/golang/glog"
	"github.com/mishudark/eventhus"
	"github.com/mishudark/eventhus/commandhandler/basic"
	"github.com/mishudark/eventhus/config"
	"github.com/mishudark/eventhus/utils"
	"os"
)

func getConfig() (eventhus.CommandBus, error) {
	// register events
	reg := eventhus.NewEventRegister()
	reg.Set(domain.PaymentPerformed{})
	reg.Set(domain.PaymentRetracted{})
	reg.Set(domain.PaymentAttributesUpdated{})

	return config.NewClient(
		config.Mongo("192.168.99.100", 27017, "organisations"),                    // event store
		config.Nats("nats://ruser:T0pS3cr3t@localhost:4222", false), // event bus
		config.AsyncCommandBus(30),                                  // command bus
		config.WireCommands(
			&domain.Payment{},        // aggregate
			basic.NewCommandHandler,  // command handler
			"organisations",   // event store bucket
			"payment",        // event store subset
			domain.PerformPayment{}, // command
			domain.UpdatePayment{},  // command
			domain.DeletePayment{},  // command
		),
	)
}

func main() {
	flag.Parse()

	commandBus, err := getConfig()
	if err != nil {
		glog.Infoln(err)
		os.Exit(1)
	}

	end := make(chan bool)

	// Add Sample Payments
	for i := 0; i < 3; i++ {
		go func() {
			uuid, err := utils.UUID()
			if err != nil {
				return
			}

			// 1) Create an account
			// var newPayment domain.PerformPayment
			// newPayment.AggregateID = uuid
			// newPayment.OrganizationId = "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb"
			// newPayment.Attributes = domain.Attributes{
			// 	Amount: 100.21,
			// 	// beneficiaryParty     BeneficiaryParty
			// 	// chargesInformation   ChargesInformation
			// 	// currency             string
			// 	// debtorParty          DebtorParty
			// 	// endToEndReference    string
			// 	// fx                   Fx
			// 	// numericReference     int
			// 	// paymentId            string
			// 	// paymentPurpose       string
			// 	// paymentScheme        string
			// 	// paymentType          string
			// 	// processingDate       string
			// 	// reference            string
			// 	// schemePaymentSubType string
			// 	// schemePaymentType    string
			// 	// sponsorParty         SponsorParty
			// }
			//
			// commandBus.HandleCommand(newPayment)
			// glog.Infof("Payment %s - payment performed", uuid)

			// 2) UpdatePayment
			// time.Sleep(time.Millisecond * 100)
			updatePaymentAttributes := domain.UpdatePayment{
				Attributes: domain.Attributes{
					// beneficiaryParty     BeneficiaryParty
					// chargesInformation   ChargesInformation
					// currency             string
					// debtorParty          DebtorParty
					// endToEndReference    string
					// fx                   Fx
					// numericReference     int
					// paymentId            string
					PaymentPurpose: "Paying for goods/services",
					PaymentScheme: "FPS",
					PaymentType: "Credit",
					// processingDate       string
					// reference            string
					// schemePaymentSubType string
					// schemePaymentType    string
					// sponsorParty         SponsorParty
				},
			}
			updatePaymentAttributes.AggregateID = uuid
			updatePaymentAttributes.Version = 1

			commandBus.HandleCommand(updatePaymentAttributes)
			glog.Infof("Payment %s - update Payment Attributes performed", uuid)

			// 3) Delete Payment
			// time.Sleep(time.Millisecond * 100)
			deletePayment := domain.DeletePayment{}
			deletePayment.AggregateID = uuid
			deletePayment.Version = 2

			commandBus.HandleCommand(deletePayment)
			glog.Infof("Payment %s - removed from storage", uuid)
		}()
	}
	<-end
}
