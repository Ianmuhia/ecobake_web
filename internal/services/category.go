package services

import (
	"context"
	"ecobake/cmd/config"
	"ecobake/internal/models"
)

type categoriesService struct {
	cfg *config.AppConfig
}

func NewCategoriesService(cfg *config.AppConfig) *categoriesService {
	return &categoriesService{cfg: cfg}
}

type CategoriesService interface {
	Create(ctx context.Context, ct models.Category) error
	Update(ctx context.Context, id int32, ct models.Category) error
	Delete(ctx context.Context, id int32) error
	List(ctx context.Context) ([]*models.Category, int, error)
}

func (cs *categoriesService) Create(ctx context.Context, ct models.Category) error {
	// if err := cs.q.CreateCategory(ctx, db.CreateCategoryParams{
	// 	Icon: sql.NullString{String: ct.Icon, Valid: true},
	// 	Name: ct.Name,
	// }); err != nil {
	// 	return err
	// }
	return nil
}

func (cs *categoriesService) Update(ctx context.Context, id int32, ct models.Category) error {
	// err := cs.q.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	Name: ct.Name,
	// 	Icon: sql.NullString{
	// 		String: ct.Icon,
	// 		Valid:  true,
	// 	},
	// 	ID: int64(ct.ID),
	// })
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (cs *categoriesService) Delete(ctx context.Context, id int32) error {
	// err := cs.q.DeleteCategory(ctx, int64(id))
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (cs *categoriesService) List(ctx context.Context) ([]*models.Category, int, error) {
	// cat, err := cs.q.ListAllCategories(ctx)

	// values := make([]*models.Category, len(cat))
	// for i, v := range cat {
	// 	values[i] = &models.Category{
	// 		ID:   int(v.ID),
	// 		Name: v.Name,
	// 		Icon: fmt.Sprintf("%s/%s/%s", cs.cfg.StorageURL.String(), cs.cfg.StorageBucket, v.Icon.String),
	// 	}
	// }
	// if err != nil {
	// 	return values, 0, err
	// }
	// return values, len(values), nil
	return []*models.Category{}, 0, nil
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
