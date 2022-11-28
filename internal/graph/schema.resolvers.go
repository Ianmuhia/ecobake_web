package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecobake/ent"
	"ecobake/ent/category"
	"ecobake/ent/favourites"
	"ecobake/ent/product"
	"ecobake/ent/user"
	"ecobake/internal/graph/generated"
	"ecobake/internal/models"
	"ecobake/pkg/randomcode"
	"fmt"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.AccountRegister, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	_, err = r.UserService.CreateUser(ctx, models.User{

		Email:       input.Email,
		Password:    string(password),
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
func (r *mutationResolver) CreateCategory(ctx context.Context, input models.CreateCategory) (*models.Category, error) {
	data, err := r.Client.Category.Create().SetIcon(input.Icon).SetName(input.Name).Save(ctx)

	if err != nil {
		return nil, err
	}
	return &models.Category{
		ID:        data.ID,
		Name:      data.Name,
		Icon:      data.Icon,
		CreatedAt: data.CreatedAt.String(),
	}, nil
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, input models.CreateCategory) (*models.Category, error) {
	save, err := r.Client.Category.Update().SetName(input.Name).SetUpdatedAt(time.Now()).SetIcon(input.Icon).Save(ctx)
	if err != nil {
		return nil, err
	}
	only, err := r.Client.Category.Query().Where(category.ID(save)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &models.Category{
		ID:        only.ID,
		Name:      only.Icon,
		Icon:      only.Name,
		UpdatedAt: only.UpdatedAt.String(),
	}, nil
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, input int) (bool, error) {
	_, err := r.Client.Category.Delete().Where(category.IDIn(input)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input *models.NewProduct) (*models.ProductCreateResponse, error) {
	data, err := r.Client.Product.Create().
		SetCreatedAt(time.Now()).SetName(input.Name).
		SetDescription(input.Description).
		SetIngredients(input.Ingredients).
		SetCategoryID(input.Category).
		SetPrice(input.Price).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			return &models.ProductCreateResponse{
				Errors: []models.ProductErrorCode{
					models.ProductErrorCodeDuplicatedInputItem,
				},
				Product: models.Product{},
			}, nil
		}
		return &models.ProductCreateResponse{
			Errors: []models.ProductErrorCode{
				models.ProductErrorCodeGraphqlError,
			},
			Product: models.Product{},
		}, nil

	}

	upload, err := r.StorageService.ProductMultipleFileUpload(input.Images, data.ID)
	if err != nil {
		return &models.ProductCreateResponse{
			Errors: []models.ProductErrorCode{
				models.ProductErrorCodeImageUploadError,
			},
			Product: models.Product{},
		}, nil
	}
	_, err = r.Client.Product.Update().SetImages(upload).Save(ctx)
	if err != nil {
		return &models.ProductCreateResponse{
			Errors: []models.ProductErrorCode{
				models.ProductErrorCodeGraphqlError,
			},
			Product: models.Product{},
		}, nil
	}
	return &models.ProductCreateResponse{
		Errors: nil,
		Product: models.Product{
			ID:          data.ID,
			Name:        data.Name,
			Price:       data.Price,
			Description: data.Description,
			Ingredients: data.Ingredients,
			TotalRating: data.TotalRating,
			Images:      nil,
			CreatedAt:   data.CreatedAt.String(),
		},
	}, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input *models.UpdateProduct) (*models.ProductCreateResponse, error) {
	data, err := r.Client.Product.UpdateOneID(input.ID).
		SetCreatedAt(time.Now()).SetName(input.Name).
		SetDescription(input.Description).
		SetIngredients(input.Ingredients).
		SetCategoryID(input.Category).
		SetPrice(input.Price).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			return &models.ProductCreateResponse{
				Errors: []models.ProductErrorCode{
					models.ProductErrorCodeDuplicatedInputItem,
				},
				Product: models.Product{},
			}, nil
		}
		return &models.ProductCreateResponse{
			Errors: []models.ProductErrorCode{
				models.ProductErrorCodeGraphqlError,
			},
			Product: models.Product{},
		}, nil

	}

	upload, err := r.StorageService.ProductMultipleFileUpload(input.Images, data.ID)
	if err != nil {
		return &models.ProductCreateResponse{
			Errors: []models.ProductErrorCode{
				models.ProductErrorCodeImageUploadError,
			},
			Product: models.Product{},
		}, nil
	}
	_, err = r.Client.Product.Update().SetImages(upload).Save(ctx)
	if err != nil {
		return &models.ProductCreateResponse{
			Errors: []models.ProductErrorCode{
				models.ProductErrorCodeGraphqlError,
			},
			Product: models.Product{},
		}, nil
	}
	return &models.ProductCreateResponse{
		Errors: nil,
		Product: models.Product{
			ID:          data.ID,
			Name:        data.Name,
			Price:       data.Price,
			Description: data.Description,
			Ingredients: data.Ingredients,
			TotalRating: data.TotalRating,
			Images:      nil,
			CreatedAt:   data.CreatedAt.String(),
		},
	}, nil
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, input int) (bool, error) {
	err := r.Client.Product.DeleteOneID(input).Exec(ctx)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// DeleteFavourite is the resolver for the deleteFavourite field.
func (r *mutationResolver) DeleteFavourite(ctx context.Context, input int) (bool, error) {
	err := r.Client.Favourites.DeleteOneID(input).Exec(ctx)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// AddFavourite is the resolver for the addFavourite field.
func (r *mutationResolver) AddFavourite(ctx context.Context, input int) (*models.ProductResponse, error) {
	data, err := r.Client.Product.Get(ctx, input)
	if err != nil {
		if ent.IsNotFound(err) {
			return &models.ProductResponse{
				Products: nil,
				Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
			}, nil
		}
		return &models.ProductResponse{
			Products: nil,
			Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeGraphqlError},
		}, nil
	}
	_, err = r.Client.Favourites.Create().SetCreatedAt(time.Now()).SetProductID(data.ID).SetUserID(1).Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) || ent.IsValidationError(err) {
			return &models.ProductResponse{
				Products: nil,
				Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeGraphqlError},
			}, nil
		}
		return &models.ProductResponse{
			Products: nil,
			Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeGraphqlError},
		}, nil
	}
	return nil, nil
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

	ok := randomcode.CheckPasswordHash(password, user.Password)
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
	accessToken, err := r.TokenService.CreateToken(user.PhoneNumber, user.Email, duration, rtduration, int64(user.ID))
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
		User: &models.User{
			ID:           user.ID,
			Email:        user.Email,
			PhoneNumber:  user.PhoneNumber,
			ProfileImage: user.ProfileImage,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		},
		Errors: nil,
	}, nil
}

// TokenRefresh is the resolver for the tokenRefresh field.
func (r *mutationResolver) TokenRefresh(ctx context.Context, refreshToken string) (*models.RefreshToken, error) {
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
	payload, err := r.TokenService.VerifyToken(token)
	if err != nil {
		return &models.VerifyToken{
			Errors: []*models.AccountError{
				{
					Field:   "refresh token ",
					Message: "provided token has expired",
					Code:    models.AccountErrorCodeJwtInvalidToken,
				},
			},
		}, nil
	}
	user, err := r.UserService.GetUserByID(ctx, payload.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			return &models.VerifyToken{
				User:    nil,
				IsValid: false,
				Payload: nil,
				Errors: []*models.AccountError{{
					Field:   "user",
					Message: "user not found",
					Code:    models.AccountErrorCodeNotFound,
				}}}, nil
		}
		return &models.VerifyToken{
			User:    nil,
			IsValid: false,
			Payload: nil,
			Errors: []*models.AccountError{{
				Field:   "user",
				Message: err.Error(),
				Code:    models.AccountErrorCodeGraphqlError,
			}},
		}, nil
	}
	return &models.VerifyToken{
		User: &models.User{
			ID:           int(user.ID),
			CreatedAt:    user.CreatedAt,
			PhoneNumber:  user.PhoneNumber,
			ProfileImage: user.ProfileImage,
			Email:        user.Email,
		},
		IsValid: true,
		Payload: nil,
		Errors:  nil,
	}, nil
}

// TokensDeactivateAll is the resolver for the tokensDeactivateAll field.
func (r *mutationResolver) TokensDeactivateAll(ctx context.Context) (*models.DeactivateAllUserTokens, error) {
	panic(fmt.Errorf("not implemented: TokensDeactivateAll - tokensDeactivateAll"))
}

// RequestPasswordReset is the resolver for the requestPasswordReset field.
func (r *mutationResolver) RequestPasswordReset(ctx context.Context, email string) (*models.RequestPasswordReset, error) {
	// user, err := r.UserService.GetUserByEmail(ctx, email)
	// if err != nil {
	// 	return &models.RequestPasswordReset{Errors: []*models.AccountError{{
	// 		Field:   "user",
	// 		Message: "provide a valid email or create account to continue",
	// 		Code:    models.AccountErrorCodeNotFound,
	// 	}}}, nil
	// }

	// // store the password reset code to db
	// verificationData := &services.VerificationData{
	// 	Email:     user.Email,
	// 	Code:      mailData.Code,
	// 	Type:      string(rune(services.PassReset)),
	// 	ExpiresAt: time.Now().Add(time.Minute * time.Duration(10)),
	// }

	// var b bytes.Buffer
	// if err := gob.NewEncoder(&b).Encode(verificationData); err != nil {
	// 	return &models.RequestPasswordReset{
	// 		Errors: []*models.AccountError{{
	// 			Field:   "password.Reset",
	// 			Message: "Unable to send password reset code. Please try again later",
	// 			Code:    models.AccountErrorCodeGraphqlError,
	// 		}},
	// 		NatsErrors: "",
	// 	}, nil
	// }

	//err = r.app.RedisConn.Set(ctx, verificationData.Email, b.Bytes(), time.Minute*time.Duration(r.app.PasswordResetCodeExpiry)).Err()
	//if err != nil {
	//	log.Println(err)
	//	restErr := resterrors.NewBadRequestError("Unable to send password reset code. Please try again later")
	//	ctx.AbortWithStatusJSON(restErr.Status, restErr)
	//	return
	//}

	//return nil, nil
	panic(fmt.Errorf("not implemented: AccountUpdate - accountUpdate"))
}

// ConfirmAccount is the resolver for the confirmAccount field.
func (r *mutationResolver) ConfirmAccount(ctx context.Context, email string, otp string) (*models.ConfirmAccount, error) {
	_, err := r.UserService.UpdateUserStatus(ctx, email)
	if err != nil {
		return &models.ConfirmAccount{Errors: []*models.AccountError{{
			Field:   "user",
			Message: "unable to update user status",
			Code:    models.AccountErrorCodeGraphqlError,
		}}}, nil
	}
	return nil, nil
}

// SetPassword is the resolver for the setPassword field.
func (r *mutationResolver) SetPassword(ctx context.Context, email string, password string, token string) (*models.SetPassword, error) {
	panic(fmt.Errorf("not implemented: AccountUpdate - accountUpdate"))
}

// PasswordChange is the resolver for the passwordChange field.
func (r *mutationResolver) PasswordChange(ctx context.Context, newPassword string, oldPassword string, accessToken string) (*models.PasswordChange, error) {
	payload, err := r.TokenService.VerifyToken(accessToken)
	if err != nil {
		return &models.PasswordChange{
			Errors: []*models.AccountError{
				{
					Field:   "refresh token ",
					Message: "provided token has expired",
					Code:    models.AccountErrorCodeJwtInvalidToken,
				},
			},
		}, nil
	}
	u, err := r.UserService.GetUserByID(ctx, payload.ID)
	if err != nil {
		return &models.PasswordChange{
			Errors: []*models.AccountError{
				{
					Field:   "refresh token ",
					Message: "provided token has expired",
					Code:    models.AccountErrorCodeJwtInvalidToken,
				},
			},
		}, nil
	}
	ok := randomcode.CheckPasswordHash(oldPassword, u.Password)
	if !ok {
		return &models.PasswordChange{
			User: nil,
			Errors: []*models.AccountError{{
				Field:   "user.Password",
				Message: "provided password does not match",
				Code:    models.AccountErrorCodeInvalidPassword,
			}},
		}, nil
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return &models.PasswordChange{
			User: nil,
			Errors: []*models.AccountError{{
				Field:   "user.Password",
				Message: "error occurred",
				Code:    models.AccountErrorCodeGraphqlError,
			}},
		}, nil
	}

	err = r.UserService.UpdateUserPassword(ctx, payload.ID, string(passwordHash))
	if err != nil {

		return &models.PasswordChange{
			User: nil,
			Errors: []*models.AccountError{{
				Field:   "user.Password",
				Message: "error occurred",
				Code:    models.AccountErrorCodeGraphqlError,
			}},
		}, nil
	}
	return &models.PasswordChange{
		User:   nil,
		Errors: nil,
	}, nil
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
	_, err := r.Client.User.Update().SetUpdatedAt(time.Now()).SetUserName(input.FirstName).Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &models.AccountUpdate{
				Errors: []*models.AccountError{
					{
						Field:   "models.User",
						Message: "User Account Not Found",
						Code:    models.AccountErrorCodeNotFound,
					},
				},
				User: nil,
			}, err
		}
		return &models.AccountUpdate{
			Errors: []*models.AccountError{
				{
					Field:   "models.User",
					Message: "Error",
					Code:    models.AccountErrorCodeGraphqlError,
				},
			},
			User: nil,
		}, nil
	}
	return nil, nil
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
func (r *mutationResolver) UserAvatarDelete(ctx context.Context, token string) (*models.UserAvatarDelete, error) {
	panic(fmt.Errorf("not implemented: UserAvatarDelete - userAvatarDelete"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) (models.Users, error) {
	users, err := r.Client.User.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return models.Users{
				Users:  nil,
				Errors: []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
			}, nil
		}
		return models.Users{
			Users:  nil,
			Errors: nil,
		}, err
	}
	m := make([]*models.User, len(users))
	for i, user := range users {
		m[i] = &models.User{
			ID:           user.ID,
			Email:        user.Email,
			PhoneNumber:  user.PhoneNumber,
			ProfileImage: user.ProfileImage,
			CreatedAt:    user.CreatedAt.String(),
		}

	}
	return models.Users{
		Users:  m,
		Errors: nil,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) (*models.Categories, error) {
	data, err := r.Client.Category.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &models.Categories{
				Categories: nil,
				Errors: []models.ListEntityErrorCode{
					models.ListEntityErrorCodeNotFound,
				},
			}, nil
		}
		return &models.Categories{
			Categories: nil,
			Errors: []models.ListEntityErrorCode{
				models.ListEntityErrorCodeGraphqlError,
			},
		}, nil
		//return nil,
	}
	m := make([]*models.Category, len(data))
	for i, datum := range data {
		m[i] = &models.Category{
			ID:        datum.ID,
			Name:      datum.Name,
			Icon:      datum.Icon,
			CreatedAt: datum.CreatedAt.String(),
		}
	}
	return &models.Categories{Categories: m}, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) (*models.Products, error) {
	data, err := r.Client.Product.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &models.Products{
				Products: nil,
				Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
			}, nil
		}
		return &models.Products{
			Products: nil,
			Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeGraphqlError},
		}, nil
	}
	pds := make([]*models.Product, len(data))

	for i, v := range data {
		pds[i] = &models.Product{
			ID:          v.ID,
			Name:        v.Name,
			Price:       v.Price,
			Description: v.Description,
			Ingredients: v.Ingredients,
			TotalRating: v.TotalRating,
			Images:      v.Images,
			CreatedAt:   v.CreatedAt.String(),
		}
	}

	return &models.Products{
		Products: pds,
		Errors:   nil,
	}, nil
}

// FavouriteProducts is the resolver for the favouriteProducts field.
func (r *queryResolver) FavouriteProducts(ctx context.Context) (*models.Products, error) {
	//TODO add actual user here
	data, err := r.Client.Favourites.Query().Where(favourites.HasUserWith(user.ID(1))).QueryProduct().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &models.Products{
				Products: nil,
				Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
			}, nil
		}
		return &models.Products{
			Products: nil,
			Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeGraphqlError},
		}, nil
	}
	pds := make([]*models.Product, len(data))

	for i, v := range data {
		pds[i] = &models.Product{
			ID:          v.ID,
			Name:        v.Name,
			Price:       v.Price,
			Description: v.Description,
			Ingredients: v.Ingredients,
			TotalRating: v.TotalRating,
			Images:      v.Images,
			CreatedAt:   v.CreatedAt.String(),
		}
	}

	return &models.Products{
		Products: pds,
		Errors:   nil,
	}, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, input int) (*models.ProductResponse, error) {
	v, err := r.Client.Product.Get(ctx, input)
	if err != nil {
		if ent.IsNotFound(err) {
			return &models.ProductResponse{
				Products: nil,
				Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
			}, nil
		}
		return &models.ProductResponse{
			Products: nil,
			Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeGraphqlError},
		}, nil
	}
	return &models.ProductResponse{
		Products: &models.Product{
			ID:          v.ID,
			Name:        v.Name,
			Price:       v.Price,
			Description: v.Description,
			Ingredients: v.Ingredients,
			TotalRating: v.TotalRating,
			Images:      v.Images,
			CreatedAt:   v.CreatedAt.String(),
		},
	}, nil
}

// ProductByCategory is the resolver for the productByCategory field.
func (r *queryResolver) ProductByCategory(ctx context.Context, input int) (*models.Products, error) {
	data, err := r.Client.Product.Query().Where(product.HasCategoryWith(category.ID(input))).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &models.Products{
				Products: nil,
				Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeNotFound},
			}, nil
		}
		return &models.Products{
			Products: nil,
			Errors:   []models.ListEntityErrorCode{models.ListEntityErrorCodeGraphqlError},
		}, nil
	}
	pds := make([]*models.Product, len(data))

	for i, v := range data {
		pds[i] = &models.Product{
			ID:          v.ID,
			Name:        v.Name,
			Price:       v.Price,
			Description: v.Description,
			Ingredients: v.Ingredients,
			TotalRating: v.TotalRating,
			Images:      v.Images,
			CreatedAt:   v.CreatedAt.String(),
		}
	}

	return &models.Products{
		Products: pds,
		Errors:   nil,
	}, nil
}

// UserCreated is the resolver for the userCreated field.
func (r *subscriptionResolver) UserCreated(ctx context.Context) (<-chan models.User, error) {
	userChan := make(<-chan models.User)

	go func() {
		v := <-r.UserChan
		log.Println(v)
	}()

	return userChan, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
