package graphql

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Attributes struct {
		Amount               func(childComplexity int) int
		BeneficiaryParty     func(childComplexity int) int
		ChargesInformation   func(childComplexity int) int
		Currency             func(childComplexity int) int
		DebtorParty          func(childComplexity int) int
		EndToEndReference    func(childComplexity int) int
		Fx                   func(childComplexity int) int
		NumericReference     func(childComplexity int) int
		PaymentId            func(childComplexity int) int
		PaymentPurpose       func(childComplexity int) int
		PaymentScheme        func(childComplexity int) int
		PaymentType          func(childComplexity int) int
		ProcessingDate       func(childComplexity int) int
		Reference            func(childComplexity int) int
		SchemePaymentSubType func(childComplexity int) int
		SchemePaymentType    func(childComplexity int) int
		SponsorParty         func(childComplexity int) int
	}

	BeneficiaryParty struct {
		AccountName       func(childComplexity int) int
		AccountNumber     func(childComplexity int) int
		AccountNumberCode func(childComplexity int) int
		AccountType       func(childComplexity int) int
		Address           func(childComplexity int) int
		BankId            func(childComplexity int) int
		BankIdCode        func(childComplexity int) int
		Name              func(childComplexity int) int
	}

	Charge struct {
		Amount   func(childComplexity int) int
		Currency func(childComplexity int) int
	}

	ChargesInformation struct {
		BearerCode              func(childComplexity int) int
		SenderCharges           func(childComplexity int) int
		ReceiverChargesAmount   func(childComplexity int) int
		ReceiverChargesCurrency func(childComplexity int) int
	}

	DebtorParty struct {
		AccountName       func(childComplexity int) int
		AccountNumber     func(childComplexity int) int
		AccountNumberCode func(childComplexity int) int
		Address           func(childComplexity int) int
		BankId            func(childComplexity int) int
		BankIdCode        func(childComplexity int) int
		Name              func(childComplexity int) int
	}

	Fx struct {
		ContractReference func(childComplexity int) int
		ExchangeRate      func(childComplexity int) int
		OriginalAmount    func(childComplexity int) int
		OriginalCurrency  func(childComplexity int) int
	}

	Mutation struct {
		StorePayment  func(childComplexity int, input StorePaymentInput) int
		RemovePayment func(childComplexity int, input RemovePaymentInput) int
	}

	PageInfo struct {
		HasPreviousPage func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		StartCursor     func(childComplexity int) int
		EndCursor       func(childComplexity int) int
	}

	Payment struct {
		Id             func(childComplexity int) int
		Type           func(childComplexity int) int
		Version        func(childComplexity int) int
		OrganisationId func(childComplexity int) int
		Attributes     func(childComplexity int) int
	}

	PaymentsConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	Query struct {
		Payment  func(childComplexity int, id string) int
		Payments func(childComplexity int, first *string, after *string, last *string, before *string, sort *Sorting) int
	}

	RemovePaymentPayload struct {
		Payment func(childComplexity int) int
	}

	SponsorParty struct {
		AccountNumber func(childComplexity int) int
		BankId        func(childComplexity int) int
		BankIdCode    func(childComplexity int) int
	}

	StorePaymentPayload struct {
		Payment func(childComplexity int) int
	}
}

type MutationResolver interface {
	StorePayment(ctx context.Context, input StorePaymentInput) (*StorePaymentPayload, error)
	RemovePayment(ctx context.Context, input RemovePaymentInput) (*RemovePaymentPayload, error)
}
type QueryResolver interface {
	Payment(ctx context.Context, id string) (*Payment, error)
	Payments(ctx context.Context, first *string, after *string, last *string, before *string, sort *Sorting) (*PaymentsConnection, error)
}

func field_Mutation_storePayment_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 StorePaymentInput
	if tmp, ok := rawArgs["input"]; ok {
		var err error
		arg0, err = UnmarshalStorePaymentInput(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil

}

func field_Mutation_removePayment_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 RemovePaymentInput
	if tmp, ok := rawArgs["input"]; ok {
		var err error
		arg0, err = UnmarshalRemovePaymentInput(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil

}

func field_Query_payment_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		var err error
		arg0, err = graphql.UnmarshalID(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil

}

func field_Query_payments_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 *string
	if tmp, ok := rawArgs["first"]; ok {
		var err error
		var ptr1 string
		if tmp != nil {
			ptr1, err = graphql.UnmarshalString(tmp)
			arg0 = &ptr1
		}

		if err != nil {
			return nil, err
		}
	}
	args["first"] = arg0
	var arg1 *string
	if tmp, ok := rawArgs["after"]; ok {
		var err error
		var ptr1 string
		if tmp != nil {
			ptr1, err = graphql.UnmarshalString(tmp)
			arg1 = &ptr1
		}

		if err != nil {
			return nil, err
		}
	}
	args["after"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["last"]; ok {
		var err error
		var ptr1 string
		if tmp != nil {
			ptr1, err = graphql.UnmarshalString(tmp)
			arg2 = &ptr1
		}

		if err != nil {
			return nil, err
		}
	}
	args["last"] = arg2
	var arg3 *string
	if tmp, ok := rawArgs["before"]; ok {
		var err error
		var ptr1 string
		if tmp != nil {
			ptr1, err = graphql.UnmarshalString(tmp)
			arg3 = &ptr1
		}

		if err != nil {
			return nil, err
		}
	}
	args["before"] = arg3
	var arg4 *Sorting
	if tmp, ok := rawArgs["sort"]; ok {
		var err error
		var ptr1 Sorting
		if tmp != nil {
			ptr1, err = UnmarshalSorting(tmp)
			arg4 = &ptr1
		}

		if err != nil {
			return nil, err
		}
	}
	args["sort"] = arg4
	return args, nil

}

func field_Query___type_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		var err error
		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil

}

func field___Type_fields_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil

}

func field___Type_enumValues_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil

}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	switch typeName + "." + field {

	case "Attributes.amount":
		if e.complexity.Attributes.Amount == nil {
			break
		}

		return e.complexity.Attributes.Amount(childComplexity), true

	case "Attributes.beneficiary_party":
		if e.complexity.Attributes.BeneficiaryParty == nil {
			break
		}

		return e.complexity.Attributes.BeneficiaryParty(childComplexity), true

	case "Attributes.charges_information":
		if e.complexity.Attributes.ChargesInformation == nil {
			break
		}

		return e.complexity.Attributes.ChargesInformation(childComplexity), true

	case "Attributes.currency":
		if e.complexity.Attributes.Currency == nil {
			break
		}

		return e.complexity.Attributes.Currency(childComplexity), true

	case "Attributes.debtor_party":
		if e.complexity.Attributes.DebtorParty == nil {
			break
		}

		return e.complexity.Attributes.DebtorParty(childComplexity), true

	case "Attributes.end_to_end_reference":
		if e.complexity.Attributes.EndToEndReference == nil {
			break
		}

		return e.complexity.Attributes.EndToEndReference(childComplexity), true

	case "Attributes.fx":
		if e.complexity.Attributes.Fx == nil {
			break
		}

		return e.complexity.Attributes.Fx(childComplexity), true

	case "Attributes.numeric_reference":
		if e.complexity.Attributes.NumericReference == nil {
			break
		}

		return e.complexity.Attributes.NumericReference(childComplexity), true

	case "Attributes.payment_id":
		if e.complexity.Attributes.PaymentId == nil {
			break
		}

		return e.complexity.Attributes.PaymentId(childComplexity), true

	case "Attributes.payment_purpose":
		if e.complexity.Attributes.PaymentPurpose == nil {
			break
		}

		return e.complexity.Attributes.PaymentPurpose(childComplexity), true

	case "Attributes.payment_scheme":
		if e.complexity.Attributes.PaymentScheme == nil {
			break
		}

		return e.complexity.Attributes.PaymentScheme(childComplexity), true

	case "Attributes.payment_type":
		if e.complexity.Attributes.PaymentType == nil {
			break
		}

		return e.complexity.Attributes.PaymentType(childComplexity), true

	case "Attributes.processing_date":
		if e.complexity.Attributes.ProcessingDate == nil {
			break
		}

		return e.complexity.Attributes.ProcessingDate(childComplexity), true

	case "Attributes.reference":
		if e.complexity.Attributes.Reference == nil {
			break
		}

		return e.complexity.Attributes.Reference(childComplexity), true

	case "Attributes.scheme_payment_sub_type":
		if e.complexity.Attributes.SchemePaymentSubType == nil {
			break
		}

		return e.complexity.Attributes.SchemePaymentSubType(childComplexity), true

	case "Attributes.scheme_payment_type":
		if e.complexity.Attributes.SchemePaymentType == nil {
			break
		}

		return e.complexity.Attributes.SchemePaymentType(childComplexity), true

	case "Attributes.sponsor_party":
		if e.complexity.Attributes.SponsorParty == nil {
			break
		}

		return e.complexity.Attributes.SponsorParty(childComplexity), true

	case "BeneficiaryParty.account_name":
		if e.complexity.BeneficiaryParty.AccountName == nil {
			break
		}

		return e.complexity.BeneficiaryParty.AccountName(childComplexity), true

	case "BeneficiaryParty.account_number":
		if e.complexity.BeneficiaryParty.AccountNumber == nil {
			break
		}

		return e.complexity.BeneficiaryParty.AccountNumber(childComplexity), true

	case "BeneficiaryParty.account_number_code":
		if e.complexity.BeneficiaryParty.AccountNumberCode == nil {
			break
		}

		return e.complexity.BeneficiaryParty.AccountNumberCode(childComplexity), true

	case "BeneficiaryParty.account_type":
		if e.complexity.BeneficiaryParty.AccountType == nil {
			break
		}

		return e.complexity.BeneficiaryParty.AccountType(childComplexity), true

	case "BeneficiaryParty.address":
		if e.complexity.BeneficiaryParty.Address == nil {
			break
		}

		return e.complexity.BeneficiaryParty.Address(childComplexity), true

	case "BeneficiaryParty.bank_id":
		if e.complexity.BeneficiaryParty.BankId == nil {
			break
		}

		return e.complexity.BeneficiaryParty.BankId(childComplexity), true

	case "BeneficiaryParty.bank_id_code":
		if e.complexity.BeneficiaryParty.BankIdCode == nil {
			break
		}

		return e.complexity.BeneficiaryParty.BankIdCode(childComplexity), true

	case "BeneficiaryParty.name":
		if e.complexity.BeneficiaryParty.Name == nil {
			break
		}

		return e.complexity.BeneficiaryParty.Name(childComplexity), true

	case "Charge.amount":
		if e.complexity.Charge.Amount == nil {
			break
		}

		return e.complexity.Charge.Amount(childComplexity), true

	case "Charge.currency":
		if e.complexity.Charge.Currency == nil {
			break
		}

		return e.complexity.Charge.Currency(childComplexity), true

	case "ChargesInformation.bearer_code":
		if e.complexity.ChargesInformation.BearerCode == nil {
			break
		}

		return e.complexity.ChargesInformation.BearerCode(childComplexity), true

	case "ChargesInformation.sender_charges":
		if e.complexity.ChargesInformation.SenderCharges == nil {
			break
		}

		return e.complexity.ChargesInformation.SenderCharges(childComplexity), true

	case "ChargesInformation.receiver_charges_amount":
		if e.complexity.ChargesInformation.ReceiverChargesAmount == nil {
			break
		}

		return e.complexity.ChargesInformation.ReceiverChargesAmount(childComplexity), true

	case "ChargesInformation.receiver_charges_currency":
		if e.complexity.ChargesInformation.ReceiverChargesCurrency == nil {
			break
		}

		return e.complexity.ChargesInformation.ReceiverChargesCurrency(childComplexity), true

	case "DebtorParty.account_name":
		if e.complexity.DebtorParty.AccountName == nil {
			break
		}

		return e.complexity.DebtorParty.AccountName(childComplexity), true

	case "DebtorParty.account_number":
		if e.complexity.DebtorParty.AccountNumber == nil {
			break
		}

		return e.complexity.DebtorParty.AccountNumber(childComplexity), true

	case "DebtorParty.account_number_code":
		if e.complexity.DebtorParty.AccountNumberCode == nil {
			break
		}

		return e.complexity.DebtorParty.AccountNumberCode(childComplexity), true

	case "DebtorParty.address":
		if e.complexity.DebtorParty.Address == nil {
			break
		}

		return e.complexity.DebtorParty.Address(childComplexity), true

	case "DebtorParty.bank_id":
		if e.complexity.DebtorParty.BankId == nil {
			break
		}

		return e.complexity.DebtorParty.BankId(childComplexity), true

	case "DebtorParty.bank_id_code":
		if e.complexity.DebtorParty.BankIdCode == nil {
			break
		}

		return e.complexity.DebtorParty.BankIdCode(childComplexity), true

	case "DebtorParty.name":
		if e.complexity.DebtorParty.Name == nil {
			break
		}

		return e.complexity.DebtorParty.Name(childComplexity), true

	case "Fx.contract_reference":
		if e.complexity.Fx.ContractReference == nil {
			break
		}

		return e.complexity.Fx.ContractReference(childComplexity), true

	case "Fx.exchange_rate":
		if e.complexity.Fx.ExchangeRate == nil {
			break
		}

		return e.complexity.Fx.ExchangeRate(childComplexity), true

	case "Fx.original_amount":
		if e.complexity.Fx.OriginalAmount == nil {
			break
		}

		return e.complexity.Fx.OriginalAmount(childComplexity), true

	case "Fx.original_currency":
		if e.complexity.Fx.OriginalCurrency == nil {
			break
		}

		return e.complexity.Fx.OriginalCurrency(childComplexity), true

	case "Mutation.storePayment":
		if e.complexity.Mutation.StorePayment == nil {
			break
		}

		args, err := field_Mutation_storePayment_args(rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.StorePayment(childComplexity, args["input"].(StorePaymentInput)), true

	case "Mutation.removePayment":
		if e.complexity.Mutation.RemovePayment == nil {
			break
		}

		args, err := field_Mutation_removePayment_args(rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RemovePayment(childComplexity, args["input"].(RemovePaymentInput)), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "Payment.id":
		if e.complexity.Payment.Id == nil {
			break
		}

		return e.complexity.Payment.Id(childComplexity), true

	case "Payment.type":
		if e.complexity.Payment.Type == nil {
			break
		}

		return e.complexity.Payment.Type(childComplexity), true

	case "Payment.version":
		if e.complexity.Payment.Version == nil {
			break
		}

		return e.complexity.Payment.Version(childComplexity), true

	case "Payment.organisation_id":
		if e.complexity.Payment.OrganisationId == nil {
			break
		}

		return e.complexity.Payment.OrganisationId(childComplexity), true

	case "Payment.attributes":
		if e.complexity.Payment.Attributes == nil {
			break
		}

		return e.complexity.Payment.Attributes(childComplexity), true

	case "PaymentsConnection.edges":
		if e.complexity.PaymentsConnection.Edges == nil {
			break
		}

		return e.complexity.PaymentsConnection.Edges(childComplexity), true

	case "PaymentsConnection.pageInfo":
		if e.complexity.PaymentsConnection.PageInfo == nil {
			break
		}

		return e.complexity.PaymentsConnection.PageInfo(childComplexity), true

	case "PaymentsConnection.totalCount":
		if e.complexity.PaymentsConnection.TotalCount == nil {
			break
		}

		return e.complexity.PaymentsConnection.TotalCount(childComplexity), true

	case "Query.payment":
		if e.complexity.Query.Payment == nil {
			break
		}

		args, err := field_Query_payment_args(rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Payment(childComplexity, args["id"].(string)), true

	case "Query.payments":
		if e.complexity.Query.Payments == nil {
			break
		}

		args, err := field_Query_payments_args(rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Payments(childComplexity, args["first"].(*string), args["after"].(*string), args["last"].(*string), args["before"].(*string), args["sort"].(*Sorting)), true

	case "RemovePaymentPayload.payment":
		if e.complexity.RemovePaymentPayload.Payment == nil {
			break
		}

		return e.complexity.RemovePaymentPayload.Payment(childComplexity), true

	case "SponsorParty.account_number":
		if e.complexity.SponsorParty.AccountNumber == nil {
			break
		}

		return e.complexity.SponsorParty.AccountNumber(childComplexity), true

	case "SponsorParty.bank_id":
		if e.complexity.SponsorParty.BankId == nil {
			break
		}

		return e.complexity.SponsorParty.BankId(childComplexity), true

	case "SponsorParty.bank_id_code":
		if e.complexity.SponsorParty.BankIdCode == nil {
			break
		}

		return e.complexity.SponsorParty.BankIdCode(childComplexity), true

	case "StorePaymentPayload.payment":
		if e.complexity.StorePaymentPayload.Payment == nil {
			break
		}

		return e.complexity.StorePaymentPayload.Payment(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Query(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:       buf,
		Errors:     ec.Errors,
		Extensions: ec.Extensions}
}

func (e *executableSchema) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Mutation(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:       buf,
		Errors:     ec.Errors,
		Extensions: ec.Extensions,
	}
}

func (e *executableSchema) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {
	return graphql.OneShot(graphql.ErrorResponse(ctx, "subscriptions are not supported"))
}

type executionContext struct {
	*graphql.RequestContext
	*executableSchema
}

var attributesImplementors = []string{"Attributes"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Attributes(ctx context.Context, sel ast.SelectionSet, obj *Attributes) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, attributesImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Attributes")
		case "amount":
			out.Values[i] = ec._Attributes_amount(ctx, field, obj)
		case "beneficiary_party":
			out.Values[i] = ec._Attributes_beneficiary_party(ctx, field, obj)
		case "charges_information":
			out.Values[i] = ec._Attributes_charges_information(ctx, field, obj)
		case "currency":
			out.Values[i] = ec._Attributes_currency(ctx, field, obj)
		case "debtor_party":
			out.Values[i] = ec._Attributes_debtor_party(ctx, field, obj)
		case "end_to_end_reference":
			out.Values[i] = ec._Attributes_end_to_end_reference(ctx, field, obj)
		case "fx":
			out.Values[i] = ec._Attributes_fx(ctx, field, obj)
		case "numeric_reference":
			out.Values[i] = ec._Attributes_numeric_reference(ctx, field, obj)
		case "payment_id":
			out.Values[i] = ec._Attributes_payment_id(ctx, field, obj)
		case "payment_purpose":
			out.Values[i] = ec._Attributes_payment_purpose(ctx, field, obj)
		case "payment_scheme":
			out.Values[i] = ec._Attributes_payment_scheme(ctx, field, obj)
		case "payment_type":
			out.Values[i] = ec._Attributes_payment_type(ctx, field, obj)
		case "processing_date":
			out.Values[i] = ec._Attributes_processing_date(ctx, field, obj)
		case "reference":
			out.Values[i] = ec._Attributes_reference(ctx, field, obj)
		case "scheme_payment_sub_type":
			out.Values[i] = ec._Attributes_scheme_payment_sub_type(ctx, field, obj)
		case "scheme_payment_type":
			out.Values[i] = ec._Attributes_scheme_payment_type(ctx, field, obj)
		case "sponsor_party":
			out.Values[i] = ec._Attributes_sponsor_party(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_amount(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Amount, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_beneficiary_party(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.BeneficiaryParty, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*BeneficiaryParty)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._BeneficiaryParty(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_charges_information(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ChargesInformation, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ChargesInformation)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._ChargesInformation(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_currency(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Currency, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_debtor_party(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DebtorParty, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*DebtorParty)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._DebtorParty(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_end_to_end_reference(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EndToEndReference, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_fx(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fx, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Fx)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._Fx(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_numeric_reference(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NumericReference, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*int)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalInt(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_payment_id(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PaymentID, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_payment_purpose(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PaymentPurpose, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_payment_scheme(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PaymentScheme, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_payment_type(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PaymentType, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_processing_date(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ProcessingDate, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_reference(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Reference, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_scheme_payment_sub_type(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SchemePaymentSubType, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_scheme_payment_type(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SchemePaymentType, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Attributes_sponsor_party(ctx context.Context, field graphql.CollectedField, obj *Attributes) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Attributes",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SponsorParty, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*SponsorParty)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._SponsorParty(ctx, field.Selections, res)
}

var beneficiaryPartyImplementors = []string{"BeneficiaryParty"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _BeneficiaryParty(ctx context.Context, sel ast.SelectionSet, obj *BeneficiaryParty) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, beneficiaryPartyImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("BeneficiaryParty")
		case "account_name":
			out.Values[i] = ec._BeneficiaryParty_account_name(ctx, field, obj)
		case "account_number":
			out.Values[i] = ec._BeneficiaryParty_account_number(ctx, field, obj)
		case "account_number_code":
			out.Values[i] = ec._BeneficiaryParty_account_number_code(ctx, field, obj)
		case "account_type":
			out.Values[i] = ec._BeneficiaryParty_account_type(ctx, field, obj)
		case "address":
			out.Values[i] = ec._BeneficiaryParty_address(ctx, field, obj)
		case "bank_id":
			out.Values[i] = ec._BeneficiaryParty_bank_id(ctx, field, obj)
		case "bank_id_code":
			out.Values[i] = ec._BeneficiaryParty_bank_id_code(ctx, field, obj)
		case "name":
			out.Values[i] = ec._BeneficiaryParty_name(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _BeneficiaryParty_account_name(ctx context.Context, field graphql.CollectedField, obj *BeneficiaryParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "BeneficiaryParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AccountName, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _BeneficiaryParty_account_number(ctx context.Context, field graphql.CollectedField, obj *BeneficiaryParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "BeneficiaryParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AccountNumber, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _BeneficiaryParty_account_number_code(ctx context.Context, field graphql.CollectedField, obj *BeneficiaryParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "BeneficiaryParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AccountNumberCode, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _BeneficiaryParty_account_type(ctx context.Context, field graphql.CollectedField, obj *BeneficiaryParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "BeneficiaryParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AccountType, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _BeneficiaryParty_address(ctx context.Context, field graphql.CollectedField, obj *BeneficiaryParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "BeneficiaryParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Address, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _BeneficiaryParty_bank_id(ctx context.Context, field graphql.CollectedField, obj *BeneficiaryParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "BeneficiaryParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.BankID, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _BeneficiaryParty_bank_id_code(ctx context.Context, field graphql.CollectedField, obj *BeneficiaryParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "BeneficiaryParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.BankIDCode, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _BeneficiaryParty_name(ctx context.Context, field graphql.CollectedField, obj *BeneficiaryParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "BeneficiaryParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var chargeImplementors = []string{"Charge"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Charge(ctx context.Context, sel ast.SelectionSet, obj *Charge) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, chargeImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Charge")
		case "amount":
			out.Values[i] = ec._Charge_amount(ctx, field, obj)
		case "currency":
			out.Values[i] = ec._Charge_currency(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Charge_amount(ctx context.Context, field graphql.CollectedField, obj *Charge) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Charge",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Amount, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Charge_currency(ctx context.Context, field graphql.CollectedField, obj *Charge) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Charge",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Currency, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var chargesInformationImplementors = []string{"ChargesInformation"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _ChargesInformation(ctx context.Context, sel ast.SelectionSet, obj *ChargesInformation) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, chargesInformationImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ChargesInformation")
		case "bearer_code":
			out.Values[i] = ec._ChargesInformation_bearer_code(ctx, field, obj)
		case "sender_charges":
			out.Values[i] = ec._ChargesInformation_sender_charges(ctx, field, obj)
		case "receiver_charges_amount":
			out.Values[i] = ec._ChargesInformation_receiver_charges_amount(ctx, field, obj)
		case "receiver_charges_currency":
			out.Values[i] = ec._ChargesInformation_receiver_charges_currency(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _ChargesInformation_bearer_code(ctx context.Context, field graphql.CollectedField, obj *ChargesInformation) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "ChargesInformation",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.BearerCode, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _ChargesInformation_sender_charges(ctx context.Context, field graphql.CollectedField, obj *ChargesInformation) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "ChargesInformation",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SenderCharges, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*Charge)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				if res[idx1] == nil {
					return graphql.Null
				}

				return ec._Charge(ctx, field.Selections, res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) _ChargesInformation_receiver_charges_amount(ctx context.Context, field graphql.CollectedField, obj *ChargesInformation) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "ChargesInformation",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ReceiverChargesAmount, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _ChargesInformation_receiver_charges_currency(ctx context.Context, field graphql.CollectedField, obj *ChargesInformation) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "ChargesInformation",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ReceiverChargesCurrency, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var debtorPartyImplementors = []string{"DebtorParty"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _DebtorParty(ctx context.Context, sel ast.SelectionSet, obj *DebtorParty) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, debtorPartyImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("DebtorParty")
		case "account_name":
			out.Values[i] = ec._DebtorParty_account_name(ctx, field, obj)
		case "account_number":
			out.Values[i] = ec._DebtorParty_account_number(ctx, field, obj)
		case "account_number_code":
			out.Values[i] = ec._DebtorParty_account_number_code(ctx, field, obj)
		case "address":
			out.Values[i] = ec._DebtorParty_address(ctx, field, obj)
		case "bank_id":
			out.Values[i] = ec._DebtorParty_bank_id(ctx, field, obj)
		case "bank_id_code":
			out.Values[i] = ec._DebtorParty_bank_id_code(ctx, field, obj)
		case "name":
			out.Values[i] = ec._DebtorParty_name(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _DebtorParty_account_name(ctx context.Context, field graphql.CollectedField, obj *DebtorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "DebtorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AccountName, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _DebtorParty_account_number(ctx context.Context, field graphql.CollectedField, obj *DebtorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "DebtorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AccountNumber, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _DebtorParty_account_number_code(ctx context.Context, field graphql.CollectedField, obj *DebtorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "DebtorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AccountNumberCode, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _DebtorParty_address(ctx context.Context, field graphql.CollectedField, obj *DebtorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "DebtorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Address, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _DebtorParty_bank_id(ctx context.Context, field graphql.CollectedField, obj *DebtorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "DebtorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.BankID, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _DebtorParty_bank_id_code(ctx context.Context, field graphql.CollectedField, obj *DebtorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "DebtorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.BankIDCode, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _DebtorParty_name(ctx context.Context, field graphql.CollectedField, obj *DebtorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "DebtorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var fxImplementors = []string{"Fx"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Fx(ctx context.Context, sel ast.SelectionSet, obj *Fx) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, fxImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Fx")
		case "contract_reference":
			out.Values[i] = ec._Fx_contract_reference(ctx, field, obj)
		case "exchange_rate":
			out.Values[i] = ec._Fx_exchange_rate(ctx, field, obj)
		case "original_amount":
			out.Values[i] = ec._Fx_original_amount(ctx, field, obj)
		case "original_currency":
			out.Values[i] = ec._Fx_original_currency(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Fx_contract_reference(ctx context.Context, field graphql.CollectedField, obj *Fx) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Fx",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ContractReference, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Fx_exchange_rate(ctx context.Context, field graphql.CollectedField, obj *Fx) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Fx",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ExchangeRate, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*float64)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalFloat(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Fx_original_amount(ctx context.Context, field graphql.CollectedField, obj *Fx) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Fx",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OriginalAmount, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Fx_original_currency(ctx context.Context, field graphql.CollectedField, obj *Fx) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Fx",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OriginalCurrency, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var mutationImplementors = []string{"Mutation"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, mutationImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Mutation",
	})

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "storePayment":
			out.Values[i] = ec._Mutation_storePayment(ctx, field)
		case "removePayment":
			out.Values[i] = ec._Mutation_removePayment(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Mutation_storePayment(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field_Mutation_storePayment_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "Mutation",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().StorePayment(rctx, args["input"].(StorePaymentInput))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*StorePaymentPayload)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._StorePaymentPayload(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Mutation_removePayment(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field_Mutation_removePayment_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "Mutation",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().RemovePayment(rctx, args["input"].(RemovePaymentInput))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*RemovePaymentPayload)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._RemovePaymentPayload(ctx, field.Selections, res)
}

var pageInfoImplementors = []string{"PageInfo"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _PageInfo(ctx context.Context, sel ast.SelectionSet, obj *PageInfo) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, pageInfoImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PageInfo")
		case "hasPreviousPage":
			out.Values[i] = ec._PageInfo_hasPreviousPage(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "hasNextPage":
			out.Values[i] = ec._PageInfo_hasNextPage(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "startCursor":
			out.Values[i] = ec._PageInfo_startCursor(ctx, field, obj)
		case "endCursor":
			out.Values[i] = ec._PageInfo_endCursor(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _PageInfo_hasPreviousPage(ctx context.Context, field graphql.CollectedField, obj *PageInfo) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PageInfo",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.HasPreviousPage, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) _PageInfo_hasNextPage(ctx context.Context, field graphql.CollectedField, obj *PageInfo) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PageInfo",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.HasNextPage, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) _PageInfo_startCursor(ctx context.Context, field graphql.CollectedField, obj *PageInfo) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PageInfo",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.StartCursor, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _PageInfo_endCursor(ctx context.Context, field graphql.CollectedField, obj *PageInfo) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PageInfo",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EndCursor, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var paymentImplementors = []string{"Payment", "Entity"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Payment(ctx context.Context, sel ast.SelectionSet, obj *Payment) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, paymentImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Payment")
		case "id":
			out.Values[i] = ec._Payment_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "type":
			out.Values[i] = ec._Payment_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "version":
			out.Values[i] = ec._Payment_version(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "organisation_id":
			out.Values[i] = ec._Payment_organisation_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "attributes":
			out.Values[i] = ec._Payment_attributes(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Payment_id(ctx context.Context, field graphql.CollectedField, obj *Payment) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Payment",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalID(res)
}

// nolint: vetshadow
func (ec *executionContext) _Payment_type(ctx context.Context, field graphql.CollectedField, obj *Payment) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Payment",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) _Payment_version(ctx context.Context, field graphql.CollectedField, obj *Payment) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Payment",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Version, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalInt(res)
}

// nolint: vetshadow
func (ec *executionContext) _Payment_organisation_id(ctx context.Context, field graphql.CollectedField, obj *Payment) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Payment",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OrganisationID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalID(res)
}

// nolint: vetshadow
func (ec *executionContext) _Payment_attributes(ctx context.Context, field graphql.CollectedField, obj *Payment) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Payment",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Attributes, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Attributes)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._Attributes(ctx, field.Selections, res)
}

var paymentsConnectionImplementors = []string{"PaymentsConnection", "Connection"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _PaymentsConnection(ctx context.Context, sel ast.SelectionSet, obj *PaymentsConnection) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, paymentsConnectionImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PaymentsConnection")
		case "edges":
			out.Values[i] = ec._PaymentsConnection_edges(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "pageInfo":
			out.Values[i] = ec._PaymentsConnection_pageInfo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "totalCount":
			out.Values[i] = ec._PaymentsConnection_totalCount(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _PaymentsConnection_edges(ctx context.Context, field graphql.CollectedField, obj *PaymentsConnection) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PaymentsConnection",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Edges, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]Payment)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec._Payment(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) _PaymentsConnection_pageInfo(ctx context.Context, field graphql.CollectedField, obj *PaymentsConnection) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PaymentsConnection",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PageInfo, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(PageInfo)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	return ec._PageInfo(ctx, field.Selections, &res)
}

// nolint: vetshadow
func (ec *executionContext) _PaymentsConnection_totalCount(ctx context.Context, field graphql.CollectedField, obj *PaymentsConnection) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "PaymentsConnection",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TotalCount, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, queryImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
	})

	var wg sync.WaitGroup
	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "payment":
			wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				out.Values[i] = ec._Query_payment(ctx, field)
				wg.Done()
			}(i, field)
		case "payments":
			wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				out.Values[i] = ec._Query_payments(ctx, field)
				wg.Done()
			}(i, field)
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	wg.Wait()
	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Query_payment(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field_Query_payment_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Payment(rctx, args["id"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Payment)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._Payment(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_payments(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field_Query_payments_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Payments(rctx, args["first"].(*string), args["after"].(*string), args["last"].(*string), args["before"].(*string), args["sort"].(*Sorting))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*PaymentsConnection)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._PaymentsConnection(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field_Query___type_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(args["name"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Schema(ctx, field.Selections, res)
}

var removePaymentPayloadImplementors = []string{"RemovePaymentPayload"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _RemovePaymentPayload(ctx context.Context, sel ast.SelectionSet, obj *RemovePaymentPayload) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, removePaymentPayloadImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("RemovePaymentPayload")
		case "payment":
			out.Values[i] = ec._RemovePaymentPayload_payment(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _RemovePaymentPayload_payment(ctx context.Context, field graphql.CollectedField, obj *RemovePaymentPayload) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "RemovePaymentPayload",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Payment, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Payment)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._Payment(ctx, field.Selections, res)
}

var sponsorPartyImplementors = []string{"SponsorParty"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _SponsorParty(ctx context.Context, sel ast.SelectionSet, obj *SponsorParty) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, sponsorPartyImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("SponsorParty")
		case "account_number":
			out.Values[i] = ec._SponsorParty_account_number(ctx, field, obj)
		case "bank_id":
			out.Values[i] = ec._SponsorParty_bank_id(ctx, field, obj)
		case "bank_id_code":
			out.Values[i] = ec._SponsorParty_bank_id_code(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _SponsorParty_account_number(ctx context.Context, field graphql.CollectedField, obj *SponsorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "SponsorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AccountNumber, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _SponsorParty_bank_id(ctx context.Context, field graphql.CollectedField, obj *SponsorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "SponsorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.BankID, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _SponsorParty_bank_id_code(ctx context.Context, field graphql.CollectedField, obj *SponsorParty) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "SponsorParty",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.BankIDCode, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var storePaymentPayloadImplementors = []string{"StorePaymentPayload"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _StorePaymentPayload(ctx context.Context, sel ast.SelectionSet, obj *StorePaymentPayload) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, storePaymentPayloadImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("StorePaymentPayload")
		case "payment":
			out.Values[i] = ec._StorePaymentPayload_payment(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _StorePaymentPayload_payment(ctx context.Context, field graphql.CollectedField, obj *StorePaymentPayload) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "StorePaymentPayload",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Payment, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Payment)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._Payment(ctx, field.Selections, res)
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __DirectiveImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))

	for idx1 := range res {
		arr1[idx1] = func() graphql.Marshaler {
			return graphql.MarshalString(res[idx1])
		}()
	}

	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___InputValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __EnumValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __FieldImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___InputValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __InputValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __SchemaImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Type(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Directive(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __TypeImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field___Type_fields_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Field(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Type(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Type(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field___Type_enumValues_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___EnumValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___InputValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) _Connection(ctx context.Context, sel ast.SelectionSet, obj *Connection) graphql.Marshaler {
	switch obj := (*obj).(type) {
	case nil:
		return graphql.Null
	case PaymentsConnection:
		return ec._PaymentsConnection(ctx, sel, &obj)
	case *PaymentsConnection:
		return ec._PaymentsConnection(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _Entity(ctx context.Context, sel ast.SelectionSet, obj *Entity) graphql.Marshaler {
	switch obj := (*obj).(type) {
	case nil:
		return graphql.Null
	case Payment:
		return ec._Payment(ctx, sel, &obj)
	case *Payment:
		return ec._Payment(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func UnmarshalRemovePaymentInput(v interface{}) (RemovePaymentInput, error) {
	var it RemovePaymentInput
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "id":
			var err error
			it.ID, err = graphql.UnmarshalID(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func UnmarshalSorting(v interface{}) (Sorting, error) {
	var it Sorting
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "orderBy":
			var err error
			var ptr1 string
			if v != nil {
				ptr1, err = graphql.UnmarshalString(v)
				it.OrderBy = &ptr1
			}

			if err != nil {
				return it, err
			}
		case "orderDirection":
			var err error
			var ptr1 string
			if v != nil {
				ptr1, err = graphql.UnmarshalString(v)
				it.OrderDirection = &ptr1
			}

			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func UnmarshalStorePaymentInput(v interface{}) (StorePaymentInput, error) {
	var it StorePaymentInput
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "id":
			var err error
			var ptr1 string
			if v != nil {
				ptr1, err = graphql.UnmarshalID(v)
				it.ID = &ptr1
			}

			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) FieldMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver) (ret interface{}) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return res
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "schema.graphql", Input: `# Core definitions

interface Connection {
    edges: [Entity]
    pageInfo: PageInfo!
    totalCount: NaturalNumber
}

# Sort info for sorting connection's edges.
input Sorting {
    orderBy: String
    orderDirection: String
}

# Page info for paginating a connection.
type PageInfo {
    hasPreviousPage: Boolean!
    hasNextPage: Boolean!
    startCursor: Cursor
    endCursor: Cursor
}

# A cursor used in paginating connections. Represented as a base64 encoded string.
scalar Cursor

scalar OrderDirection

interface Entity {
    id: ID!
}

# A currency trading symbol e.g. USD
scalar Currency

# A natural number (integer >= 0)
scalar NaturalNumber

# A float with 2 decimals
scalar ShortFloat

# full name e.g. "Wilfred Jeremiah Owens"
scalar Name

# short version of name e.g. "W Owens"
scalar ShortName

# date of format YYYY-MM-DD e.g."2017-01-18"
scalar Date

# e.g. "10 Debtor Crescent Sourcetown NE1"
scalar Address

type Charge {
  amount: ShortFloat
  currency: Currency
}

type BeneficiaryParty {
  account_name: ShortName
  account_number: String
  account_number_code: String
  account_type: NaturalNumber
  address: Address
  bank_id: String
  bank_id_code: String
  name: Name
}

type ChargesInformation {
  bearer_code: String
  sender_charges: [Charge]
  receiver_charges_amount: ShortFloat
  receiver_charges_currency: Currency
}

type DebtorParty {
  account_name: ShortName
  account_number: String
  account_number_code: String
  address: Address
  bank_id: String
  bank_id_code: String
  name: Name
}

type Fx {
  contract_reference: String
  exchange_rate: Float
  original_amount: ShortFloat
  original_currency: Currency
}

type SponsorParty {
  account_number: String
  bank_id: String
  bank_id_code: String
}

type Attributes {
    amount: ShortFloat
    beneficiary_party: BeneficiaryParty
    charges_information: ChargesInformation
    currency: Currency
    debtor_party: DebtorParty
    end_to_end_reference: String
    fx: Fx
    numeric_reference: Int
    payment_id: String
    payment_purpose: String
    payment_scheme: String
    payment_type: String
    processing_date: Date
    reference: String
    scheme_payment_sub_type: String
    scheme_payment_type: String
    sponsor_party: SponsorParty
}

# A member of a Memorial on Memoria.
type Payment implements Entity {
    id: ID!
    type: String!
    version: Int!
    organisation_id: ID!
    attributes: Attributes
}

type PaymentsConnection implements Connection {
    edges: [Payment!]!
    pageInfo: PageInfo!
    totalCount: NaturalNumber
}

input StorePaymentInput {
    id: ID
}

type StorePaymentPayload {
    payment: Payment
}

input RemovePaymentInput {
    id: ID!
}

type RemovePaymentPayload {
    payment: Payment
}


type Query {
  # Retrieve a payment by said parameters.
  payment(
    # The ID of the payment.
    id: ID!
  ): Payment

  # Retrieve payments using Relay connection principles.
  payments(
    first: NaturalNumber = 10
    after: Cursor
    last: NaturalNumber
    before: Cursor
    sort: Sorting
  ): PaymentsConnection
}

type Mutation {
    storePayment(input: StorePaymentInput!): StorePaymentPayload
    removePayment(input: RemovePaymentInput!): RemovePaymentPayload
}
`},
)
