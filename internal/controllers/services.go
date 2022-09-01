package controllers

import (
	"ecobake/cmd/config"
	"ecobake/ent"
	"ecobake/internal/services"
)

// Repository is the repository type.
type Repository struct {
	mailService     services.MailService
	app             *config.AppConfig
	storageService  services.FileStorageService
	natService      services.NatsService
	userService     services.UsersService
	tokenService    services.TokenService
	natsService     services.NatsService
	searchService   services.SearchService
	CategoryService services.CategoriesService
	Client          *ent.Client
	//paymentService services.PaymentsService
}

// NewRepo creates a new repository.
func NewRepo(
	mailService services.MailService,
	a *config.AppConfig,
	storageService services.FileStorageService,
	natService services.NatsService,
	userService services.UsersService,
	tokenService services.TokenService,
	searchService services.SearchService,
	categoryService services.CategoriesService,
	client *ent.Client,

	//paymentService services.PaymentsService,

) *Repository {
	return &Repository{
		mailService:     mailService,
		app:             a,
		storageService:  storageService,
		natService:      natService,
		userService:     userService,
		tokenService:    tokenService,
		natsService:     natService,
		searchService:   searchService,
		CategoryService: categoryService,
		Client:          client,
		//paymentService: paymentService,
	}
}
