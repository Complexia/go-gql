package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Email is the resolver for the email field.
func (r *userResolver) Email(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented: Email - email"))
}

// Password is the resolver for the password field.
func (r *userResolver) Password(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented: Password - password"))
}

// FirstName is the resolver for the firstName field.
func (r *userResolver) FirstName(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented: FirstName - firstName"))
}

// LastName is the resolver for the lastName field.
func (r *userResolver) LastName(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented: LastName - lastName"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
