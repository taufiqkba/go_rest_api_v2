package helper

import (
	"github.com/taufiqkba/go_rest_api_v2/model/domain"
	"github.com/taufiqkba/go_rest_api_v2/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	//looping to return
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	//return all categories
	return categoryResponses
}
