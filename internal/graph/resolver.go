package graph

import (
	"ecobake/ent"
	"ecobake/internal/models"
	"ecobake/internal/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client          *ent.Client
	UserChan        chan models.User
	StorageService  services.FileStorageService
	UserService     services.UsersService
	TokenService    services.TokenService
	CategoryService services.CategoriesService
}

func NewResolver(
	userChan chan models.User,
	client *ent.Client,
	storageService services.FileStorageService,
	userService services.UsersService,
	tokenService services.TokenService,
	categoryService services.CategoriesService,

) *Resolver {
	return &Resolver{
		UserChan:        userChan,
		Client:          client,
		StorageService:  storageService,
		UserService:     userService,
		TokenService:    tokenService,
		CategoryService: categoryService,
	}
}
