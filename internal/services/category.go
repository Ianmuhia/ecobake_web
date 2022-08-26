package services

import (
	"context"
	"database/sql"
	"ecobake/cmd/config"
	"ecobake/internal/models"
	"ecobake/internal/postgresql/db"
	"fmt"
	"log"
)

type categoriesService struct {
	q   *db.Queries
	cfg *config.AppConfig
}

func NewCategoriesService(q db.DBTX, cfg *config.AppConfig) *categoriesService {
	return &categoriesService{q: db.New(q), cfg: cfg}
}

type CategoriesService interface {
	CreateCategories(ctx context.Context, ct models.Category) error
	UpdateCategory(ctx context.Context, id int32, ct models.Category) error
	DeleteCategory(ctx context.Context, id int32) error
	ListCategories(ctx context.Context) ([]*models.Category, int, error)
	//GetCategory(ctx context.Context, id int32) (*models.Category, error)
}

func (cs *categoriesService) CreateCategories(ctx context.Context, ct models.Category) error {
	if err := cs.q.CreateCategory(ctx, db.CreateCategoryParams{
		Icon: sql.NullString{String: ct.Icon, Valid: true},
		Name: ct.Name,
	}); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (cs *categoriesService) UpdateCategory(ctx context.Context, id int32, ct models.Category) error {
	err := cs.q.UpdateCategory(ctx, db.UpdateCategoryParams{
		Name: ct.Name,
		Icon: sql.NullString{
			String: ct.Icon,
			Valid:  true,
		},
		ID: int64(ct.ID),
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (cs *categoriesService) DeleteCategory(ctx context.Context, id int32) error {
	err := cs.q.DeleteCategory(ctx, int64(id))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (cs *categoriesService) ListCategories(ctx context.Context) ([]*models.Category, int, error) {
	cat, err := cs.q.ListAllCategories(ctx)

	values := make([]*models.Category, len(cat))
	for i, v := range cat {
		values[i] = &models.Category{
			ID:   int32(v.ID),
			Name: v.Name,
			Icon: fmt.Sprintf("%s/%s/%s", cs.cfg.StorageURL.String(), cs.cfg.StorageBucket, v.Icon),
		}
	}
	if err != nil {
		log.Println(err)
		return values, 0, err
	}
	return values, len(values), nil
}

//func (cs *categoriesService) GetCategory(ctx context.Context, id int32) (*models.Category, error) {
//	cat, err := cs.q.GetCategory(ctx, id)
//
//	f := &models.Category{
//		ID:        cat.ID,
//		CreatedAt: cat.CreatedAt,
//		Icon:      fmt.Sprintf("%s/%s/%s", cs.cfg.StorageURL.String(), cs.cfg.StorageBucket, cat.Icon),
//		Name:      cat.Name,
//	}
//	if err != nil {
//		log.Println(err)
//		return f, err
//	}
//	return f, err
//}
