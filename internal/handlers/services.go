package controllers

import (
	"ecobake/cmd/config"
	"ecobake/ent"
	"ecobake/internal/services"
)

// Repository is the repository type.
type Repository struct {
	app             *config.AppConfig
	storageService  services.FileStorageService
	userService     services.UsersService
	tokenService    services.TokenService
	CategoryService services.CategoriesService
	Client          *ent.Client
}

// NewRepo creates a new repository.
func NewRepo(
	a *config.AppConfig,
	storageService services.FileStorageService,
	userService services.UsersService,
	tokenService services.TokenService,
	categoryService services.CategoriesService,
	client *ent.Client,
) *Repository {
	return &Repository{
		app:             a,
		storageService:  storageService,
		userService:     userService,
		tokenService:    tokenService,
		CategoryService: categoryService,
		Client:          client,
	}
}
