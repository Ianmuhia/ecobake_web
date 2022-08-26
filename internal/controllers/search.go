package controllers

import (
	"ecobake/internal/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func (r *Repository) CreateIndexes(ctx *gin.Context) {
	var wg sync.WaitGroup
	fis := []models.FaindaIndexes{
		{
			PrimaryKey: "id",
			Uid:        "users",
			Filterableattributes: &[]string{
				"id",
				"created_at",
				"user_name",
				"email",
				"_geo",
			},
			Searchableattributes: &[]string{
				"id",
				"created_at",
				"user_name",
				"email",
				"_geo",
			},
			Sortableattributes: &[]string{
				"id",
				"created_at",
				"user_name",
				"email",
				"_geo",
			},
			Rankingrules: nil,
		},
		{
			PrimaryKey: "id",
			Uid:        "businesses",
			Filterableattributes: &[]string{"id",
				"created_at",
				"updated_at",
				"deleted_at",
				"user",
				"user_id",
				"name",
				"description",
				"location",
				"jobs_done",
				"profile_image",
				"ventures",
				"_geo",
				"experience",
				"is_visible",
				"average_rating",
				"total_ratings"},
			Searchableattributes: &[]string{
				"id",
				"created_at",
				"updated_at",
				"deleted_at",
				"user",
				"user_id",
				"name",
				"description",
				"location",
				"jobs_done",
				"profile_image",
				"ventures",
				"_geo",
				"experience",
				"is_visible",
				"average_rating",
				"total_ratings"},
			Sortableattributes: &[]string{
				"id",
				"created_at",
				"updated_at",
				"deleted_at",
				"user",
				"user_id",
				"name",
				"description",
				"location",
				"jobs_done",
				"ventures",
				"_geo",
				"experience",
				"is_visible",
				"average_rating",
				"total_ratings"},
			Rankingrules: nil,
		},
		{
			PrimaryKey: "id",
			Uid:        "ventures",
			Sortableattributes: &[]string{"id, created_at",
				"updated_at",
				"deleted_at",
				"venture_name",
				"rate",
				"business_id",
				"category_id",
				"venture_category",
				"available",
				"is_default",
				"busy"},
			Filterableattributes: &[]string{"id, created_at",
				"updated_at",
				"deleted_at",
				"venture_name",
				"rate",
				"business_id",
				"category_id",
				"venture_category",
				"available",
				"is_default",
				"busy"},
			Searchableattributes: &[]string{"id",
				"created_at",
				"updated_at",
				"deleted_at",
				"venture_name",
				"rate",
				"business_id",
				"category_id",
				"venture_category",
				"available",
				"is_default",
				"busy"},
			Rankingrules: nil,
		},
		{
			PrimaryKey:           "id",
			Uid:                  "payments",
			Filterableattributes: &[]string{"id, created_at"},
			Searchableattributes: &[]string{},
			Rankingrules:         nil,
		},
	}
	wg.Add(len(fis))
	for _, fi := range fis {
		go r.searchService.CreateIndexes(fi, &wg)
	}
	wg.Wait()

	ctx.JSON(http.StatusOK, nil)

}

func (r *Repository) SearchData(ctx *gin.Context) {

	//data, err := r.searchService.SearchUser()
	//if err != nil {
	//	log.Println(err)
	//	data := resterrors.NewBadRequestError("Error Processing request")
	//	ctx.AbortWithStatusJSON(data.Status, data)
	//	return
	//}
	//resp := NewStatusOkResponse("success", data)
	//ctx.JSON(resp.Status, resp)
}
