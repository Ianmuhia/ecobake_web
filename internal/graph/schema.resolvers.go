package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecobake/internal/graph/generated"
	"ecobake/internal/graph/model"
	"ecobake/internal/models"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (*model.LoginResp, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// UserCreated is the resolver for the userCreated field.
func (r *subscriptionResolver) UserCreated(ctx context.Context) (<-chan *models.User, error) {
	panic(fmt.Errorf("not implemented: UserCreated - userCreated"))
}

// Name is the resolver for the name field.
func (r *userResolver) Name(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: Name - name"))
}

// Workplace is the resolver for the workplace field.
func (r *userResolver) Workplace(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: Workplace - workplace"))
}

// FirstName is the resolver for the first_name field.
func (r *userResolver) FirstName(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: FirstName - first_name"))
}

// LastName is the resolver for the last_name field.
func (r *userResolver) LastName(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: LastName - last_name"))
}

// School is the resolver for the school field.
func (r *userResolver) School(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: School - school"))
}

// Phone is the resolver for the phone field.
func (r *userResolver) Phone(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: Phone - phone"))
}

// CreatedAt is the resolver for the created_at field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// UpdatedAt is the resolver for the updated_at field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updated_at"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
