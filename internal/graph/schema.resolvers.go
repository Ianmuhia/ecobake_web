package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecobake/ent"
	"ecobake/internal/graph/generated"
	"ecobake/internal/models"
	"fmt"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// Categories is the resolver for the categories field.
func (r *categoriesResolver) Categories(ctx context.Context, obj ent.Categories) ([]*ent.Category, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}

// Errors is the resolver for the errors field.
func (r *categoriesResolver) Errors(ctx context.Context, obj ent.Categories) ([]models.ListEntityErrorCode, error) {
	if len(obj) == 0 {
		return models.AllListEntityErrorCode, nil

	}
	return nil, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *categoryResolver) CreatedAt(ctx context.Context, obj *ent.Category) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *categoryResolver) UpdatedAt(ctx context.Context, obj *ent.Category) (string, error) {
	return obj.UpdatedAt.String(), nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.AccountRegister, error) {
	_, err := r.UserService.CreateUser(ctx, models.User{

		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	})
	if err != nil {
		return &models.AccountRegister{
			RequiresConfirmation: new(bool),
			Errors: []*models.AccountError{
				{
					Field:   "user.Password",
					Message: "incorrect password ",
					Code:    models.AccountErrorCodeGraphqlError,
				},
			},
			User: nil,
		}, nil
	}
	return nil, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input models.CreateCategory) (ent.Categories, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// TokenCreate is the resolver for the tokenCreate field.
func (r *mutationResolver) TokenCreate(ctx context.Context, email string, password string) (*models.CreateToken, error) {
	user, err := r.UserService.GetUserByEmail(ctx, email)
	if err != nil {
		return &models.CreateToken{
			Errors: []*models.AccountError{
				{
					Field:   "user",
					Message: "user not found",
					Code:    models.AccountErrorCodeNotFound,
				},
			},
		}, nil
	}
	m := models.User{PasswordHash: user.PasswordHash}
	ok := m.CheckPasswordHash(password)
	log.Println(ok)
	if !ok {
		return &models.CreateToken{
			Errors: []*models.AccountError{
				{
					Field:   "user.Password",
					Message: "incorrect password ",
					Code:    models.AccountErrorCodeInvalidPassword,
				},
			},
		}, nil
	}
	duration := 30 * time.Hour
	rtduration := time.Duration(time.Now().Add(time.Hour * 100).Unix())
	accessToken, err := r.TokenService.CreateToken(user.UserName, user.Email, duration, rtduration, user.ID)
	if err != nil {
		return &models.CreateToken{
			Errors: []*models.AccountError{
				{
					Field:   "token",
					Message: "unable to create token",
					Code:    models.AccountErrorCodeGraphqlError,
				},
			},
		}, nil
	}
	return &models.CreateToken{
		Token:        accessToken,
		RefreshToken: accessToken,
		User: &ent.User{
			ID:           int(user.ID),
			UserName:     user.UserName,
			CreatedAt:    user.CreatedAt,
			PhoneNumber:  user.PhoneNumber,
			IsVerified:   user.IsVerified,
			ProfileImage: user.ProfileImage,
			Email:        user.Email,
		},
		Errors: nil,
	}, err
}

// TokenRefresh is the resolver for the tokenRefresh field.
func (r *mutationResolver) TokenRefresh(_ context.Context, refreshToken string) (*models.RefreshToken, error) {
	payload, err := r.TokenService.VerifyToken(refreshToken)
	if err != nil {
		return &models.RefreshToken{
			Errors: []*models.AccountError{
				{
					Field:   "refresh token ",
					Message: "provided token has expired",
					Code:    models.AccountErrorCodeJwtInvalidToken,
				},
			},
		}, nil
	}
	duration := 30 * time.Hour
	rtduration := time.Duration(time.Now().Add(time.Hour * 100).Unix())
	accessToken, err := r.TokenService.CreateToken(payload.Username, payload.Email, duration, rtduration, payload.ID)
	if err != nil {
		return &models.RefreshToken{
			Errors: []*models.AccountError{
				{
					Field:   "token",
					Message: "unable to create token",
					Code:    models.AccountErrorCodeGraphqlError,
				},
			},
		}, nil
	}
	return &models.RefreshToken{
		Token: accessToken,
	}, nil
}

// TokenVerify is the resolver for the tokenVerify field.
func (r *mutationResolver) TokenVerify(ctx context.Context, token string) (*models.VerifyToken, error) {
	panic(fmt.Errorf("not implemented: TokenVerify - tokenVerify"))
}

// TokensDeactivateAll is the resolver for the tokensDeactivateAll field.
func (r *mutationResolver) TokensDeactivateAll(ctx context.Context) (*models.DeactivateAllUserTokens, error) {
	panic(fmt.Errorf("not implemented: TokensDeactivateAll - tokensDeactivateAll"))
}

// RequestPasswordReset is the resolver for the requestPasswordReset field.
func (r *mutationResolver) RequestPasswordReset(ctx context.Context, channel *string, email string, redirectURL string) (*models.RequestPasswordReset, error) {
	panic(fmt.Errorf("not implemented: RequestPasswordReset - requestPasswordReset"))
}

// ConfirmAccount is the resolver for the confirmAccount field.
func (r *mutationResolver) ConfirmAccount(ctx context.Context, email string, token string) (*models.ConfirmAccount, error) {
	panic(fmt.Errorf("not implemented: ConfirmAccount - confirmAccount"))
}

// SetPassword is the resolver for the setPassword field.
func (r *mutationResolver) SetPassword(ctx context.Context, email string, password string, token string) (*models.SetPassword, error) {
	panic(fmt.Errorf("not implemented: SetPassword - setPassword"))
}

// PasswordChange is the resolver for the passwordChange field.
func (r *mutationResolver) PasswordChange(ctx context.Context, newPassword string, oldPassword string) (*models.PasswordChange, error) {
	panic(fmt.Errorf("not implemented: PasswordChange - passwordChange"))
}

// RequestEmailChange is the resolver for the requestEmailChange field.
func (r *mutationResolver) RequestEmailChange(ctx context.Context, channel *string, newEmail string, password string, redirectURL string) (*models.RequestEmailChange, error) {
	panic(fmt.Errorf("not implemented: RequestEmailChange - requestEmailChange"))
}

// ConfirmEmailChange is the resolver for the confirmEmailChange field.
func (r *mutationResolver) ConfirmEmailChange(ctx context.Context, channel *string, token string) (*models.ConfirmEmailChange, error) {
	panic(fmt.Errorf("not implemented: ConfirmEmailChange - confirmEmailChange"))
}

// AccountRegister is the resolver for the accountRegister field.
func (r *mutationResolver) AccountRegister(ctx context.Context, input models.AccountRegisterInput) (*models.AccountRegister, error) {
	panic(fmt.Errorf("not implemented: AccountRegister - accountRegister"))
}

// AccountUpdate is the resolver for the accountUpdate field.
func (r *mutationResolver) AccountUpdate(ctx context.Context, input models.AccountInput) (*models.AccountUpdate, error) {
	panic(fmt.Errorf("not implemented: AccountUpdate - accountUpdate"))
}

// AccountRequestDeletion is the resolver for the accountRequestDeletion field.
func (r *mutationResolver) AccountRequestDeletion(ctx context.Context, channel *string, redirectURL string) (*models.AccountRequestDeletion, error) {
	panic(fmt.Errorf("not implemented: AccountRequestDeletion - accountRequestDeletion"))
}

// AccountDelete is the resolver for the accountDelete field.
func (r *mutationResolver) AccountDelete(ctx context.Context, token string) (*models.AccountDelete, error) {
	panic(fmt.Errorf("not implemented: AccountDelete - accountDelete"))
}

// UserAvatarUpdate is the resolver for the userAvatarUpdate field.
func (r *mutationResolver) UserAvatarUpdate(ctx context.Context, image graphql.Upload) (*models.UserAvatarUpdate, error) {
	panic(fmt.Errorf("not implemented: UserAvatarUpdate - userAvatarUpdate"))
}

// UserAvatarDelete is the resolver for the userAvatarDelete field.
func (r *mutationResolver) UserAvatarDelete(ctx context.Context) (*models.UserAvatarDelete, error) {
	panic(fmt.Errorf("not implemented: UserAvatarDelete - userAvatarDelete"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) (ent.Users, error) {
	users, err := r.Client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) (ent.Categories, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}

// UserCreated is the resolver for the userCreated field.
func (r *subscriptionResolver) UserCreated(ctx context.Context) (<-chan ent.User, error) {
	panic(fmt.Errorf("not implemented: UserCreated - userCreated"))
}

// CreatedAt is the resolver for the created_at field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *ent.User) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *ent.User) (string, error) {
	return obj.UpdatedAt.String(), nil
}

// Users is the resolver for the users field.
func (r *usersResolver) Users(ctx context.Context, obj ent.Users) ([]*ent.User, error) {
	users, err := r.Client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Errors is the resolver for the errors field.
func (r *usersResolver) Errors(ctx context.Context, obj ent.Users) ([]models.ListEntityErrorCode, error) {
	if len(obj) != 0 {
		return nil, nil
	}

	return models.AllListEntityErrorCode, nil
}

// Categories returns generated.CategoriesResolver implementation.
func (r *Resolver) Categories() generated.CategoriesResolver { return &categoriesResolver{r} }

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// Users returns generated.UsersResolver implementation.
func (r *Resolver) Users() generated.UsersResolver { return &usersResolver{r} }

type categoriesResolver struct{ *Resolver }
type categoryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type usersResolver struct{ *Resolver }
