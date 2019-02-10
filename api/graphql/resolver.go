package graphql

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) StorePayment(ctx context.Context, input StorePaymentInput) (*StorePaymentPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) RemovePayment(ctx context.Context, input RemovePaymentInput) (*RemovePaymentPayload, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Payment(ctx context.Context, id string) (*Payment, error) {
	panic("not implemented")
}
func (r *queryResolver) Payments(ctx context.Context, first *string, after *string, last *string, before *string, sort *Sorting) (*PaymentsConnection, error) {
	panic("not implemented")
}
