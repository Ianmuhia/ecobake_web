package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecobake/internal/graph/generated"
	"ecobake/internal/models"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// CreatedAt is the resolver for the created_at field.
func (r *categoryResolver) CreatedAt(ctx context.Context, obj *models.Category) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *categoryResolver) UpdatedAt(ctx context.Context, obj *models.Category) (string, error) {
	return obj.UpdatedAt.String(), nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.AccountRegister, error) {
	data := models.User{
		UserName:     "",
		Email:        input.Email,
		PasswordHash: "",
		PhoneNumber:  input.PhoneNumber,
		ProfileImage: "",
		IsVerified:   false,
	}
	//data.PasswordHash = data.Hash()
	_, err := r.UserService.CreateUser(ctx, data)
	if err != nil {
		return &models.AccountRegister{
			Errors: []*models.AccountError{},
		}, nil
	}
	return &models.AccountRegister{
		Errors: []*models.AccountError{{
			Field:   nil,
			Message: nil,
			Code:    models.AccountErrorCodeAccountNotConfirmed,
		}, {
			Field:   nil,
			Message: nil,
			Code:    models.AccountErrorCodeActivateSuperuserAccount,
		}},
		User: &models.User{
			ID:           0,
			CreatedAt:    time.Time{},
			UpdatedAt:    time.Time{},
			DeletedAt:    time.Time{},
			UserName:     "",
			Email:        "",
			PasswordHash: "",
			PhoneNumber:  "",
			Password:     "",
			ProfileImage: "",
			IsVerified:   false,
		},
	}, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input models.CreateCategory) (*models.Categories, error) {
	err := r.CategoryService.CreateCategories(ctx, models.Category{
		Name: input.Name,
		Icon: input.Name,
	})
	if err != nil {
		var pgError *pgconn.PgError
		if errors.Is(pgError, err) {

			return &models.Categories{
				Categories: nil,
				Errors:     []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
			}, nil
		}
		return &models.Categories{
			Categories: nil,
			Errors:     []models.ListEntityErrorCode{models.ListEntityErrorCodeGraphqlError},
		}, nil
	}
	return nil, nil

}

// TokenCreate is the resolver for the tokenCreate field.
func (r *mutationResolver) TokenCreate(ctx context.Context, email string, password string) (*models.CreateToken, error) {
	panic(fmt.Errorf("not implemented: TokenCreate - tokenCreate"))
}

// TokenRefresh is the resolver for the tokenRefresh field.
func (r *mutationResolver) TokenRefresh(ctx context.Context, csrfToken *string, refreshToken *string) (*models.RefreshToken, error) {
	panic(fmt.Errorf("not implemented: TokenRefresh - tokenRefresh"))
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
func (r *queryResolver) Users(ctx context.Context) (*models.Users, error) {
	_, users, err := r.UserService.GetAllUsers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &models.Users{
				Users:  nil,
				Errors: []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
			}, nil
		}
		return &models.Users{
			Users:  nil,
			Errors: []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
		}, nil
	}
	m := make([]*models.User, len(users))
	for i, user := range users {
		m[i] = &models.User{
			ID:           user.ID,
			CreatedAt:    user.CreatedAt,
			UserName:     user.UserName,
			Email:        user.Email,
			PhoneNumber:  user.PhoneNumber,
			ProfileImage: user.ProfileImage,
			IsVerified:   user.IsVerified,
		}
	}
	return &models.Users{
		Users:  m,
		Errors: nil,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) (*models.Categories, error) {
	cat, _, err := r.CategoryService.ListCategories(ctx)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return &models.Categories{
				Categories: nil,
				Errors:     []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
			}, nil
		}
		return &models.Categories{
			Categories: nil,
			Errors:     []models.ListEntityErrorCode{models.ListEntityErrorCodeGraphqlError},
		}, nil
	}
	m := make([]*models.Category, len(cat))
	for i, category := range cat {
		m[i] = &models.Category{
			ID:        category.ID,
			CreatedAt: category.CreatedAt,
			Name:      category.Name,
			Icon:      category.Icon,
		}
	}
	return &models.Categories{
		Categories: m,
		Errors:     nil,
	}, nil
}

// UserCreated is the resolver for the userCreated field.
func (r *subscriptionResolver) UserCreated(ctx context.Context) (<-chan *models.User, error) {
	panic(fmt.Errorf("not implemented: UserCreated - userCreated"))
}

// CreatedAt is the resolver for the created_at field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.UpdatedAt.String(), nil
}

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

type categoryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
