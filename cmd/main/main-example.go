package main
//
// import (
// 	"flag"
// 	"github.com/golang/glog"
// 	"github.com/mishudark/eventhus"
// 	"github.com/mishudark/eventhus/commandhandler/basic"
// 	"github.com/mishudark/eventhus/config"
// 	"github.com/mishudark/eventhus/examples/bank"
//
// 	"github.com/mishudark/eventhus/utils"
// 	"os"
// )
//
// func getConfig() (eventhus.CommandBus, error) {
// 	// register events
// 	reg := eventhus.NewEventRegister()
// 	reg.Set(bank.AccountCreated{})
// 	reg.Set(bank.DepositPerformed{})
// 	reg.Set(bank.WithdrawalPerformed{})
//
// 	// eventbus
// 	// rabbit, err := config.RabbitMq("guest", "guest", "localhost", 5672)
//
// 	return config.NewClient(
// 		config.Mongo("192.168.99.100", 27017, "bank"),                    // event store
// 		config.Nats("nats://ruser:T0pS3cr3t@localhost:4222", false), // event bus
// 		config.AsyncCommandBus(30),                                  // command bus
// 		config.WireCommands(
// 			&bank.Account{},          // aggregate
// 			basic.NewCommandHandler,  // command handler
// 			"bank",                   // event store bucket
// 			"account",                // event store subset
// 			bank.CreateAccount{},     // command
// 			bank.PerformDeposit{},    // command
// 			bank.PerformWithdrawal{}, // command
// 		),
// 	)
// }
//
// func main() {
// 	flag.Parse()
//
// 	commandBus, err := getConfig()
// 	if err != nil {
// 		glog.Infoln(err)
// 		os.Exit(1)
// 	}
//
// 	end := make(chan bool)
//
// 	// Create Account
// 	for i := 0; i < 3; i++ {
// 		go func() {
// 			uuid, err := utils.UUID()
// 			if err != nil {
// 				return
// 			}
//
// 			// 1) Create an account
// 			var account bank.CreateAccount
// 			account.AggregateID = uuid
// 			account.Owner = "mishudark"
//
// 			commandBus.HandleCommand(account)
// 			glog.Infof("account %s - account created", uuid)
//
// 			// 2) Perform a deposit
// 			// time.Sleep(time.Millisecond * 100)
// 			deposit := bank.PerformDeposit{
// 				Amount: 300,
// 			}
//
// 			deposit.AggregateID = uuid
// 			deposit.Version = 1
//
// 			commandBus.HandleCommand(deposit)
// 			glog.Infof("account %s - deposit performed", uuid)
//
// 			// 3) Perform a withdrawl
// 			// time.Sleep(time.Millisecond * 100)
// 			withdrawl := bank.PerformWithdrawal{
// 				Amount: 249,
// 			}
//
// 			withdrawl.AggregateID = uuid
// 			withdrawl.Version = 2
//
// 			commandBus.HandleCommand(withdrawl)
// 			glog.Infof("account %s - withdrawl performed", uuid)
// 		}()
// 	}
// 	<-end
// }
