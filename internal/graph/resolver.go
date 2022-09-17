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
	NatService      services.NatsService
	UserService     services.UsersService
	TokenService    services.TokenService
	SearchService   services.SearchService
	CategoryService services.CategoriesService
}

func NewResolver(
	userChan chan models.User,
	client *ent.Client,
	storageService services.FileStorageService,
	natService services.NatsService,
	userService services.UsersService,
	tokenService services.TokenService,
	searchService services.SearchService,
	categoryService services.CategoriesService,

) *Resolver {
	return &Resolver{
		UserChan:        userChan,
		Client:          client,
		StorageService:  storageService,
		NatService:      natService,
		UserService:     userService,
		TokenService:    tokenService,
		SearchService:   searchService,
		CategoryService: categoryService,
	}
}
