package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/taufiqkba/go_rest_api_v2/helper"
	"github.com/taufiqkba/go_rest_api_v2/model/domain"
	"github.com/taufiqkba/go_rest_api_v2/model/web"
	"github.com/taufiqkba/go_rest_api_v2/repository"
)

type CategoryServiceImpl struct {
	repository repository.CategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryServiceImpl(repository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{repository: repository, DB: DB, Validate: validate}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	//validate first
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	//set database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//create new category
	category := domain.Category{
		Name: request.Name,
	}
	category = service.repository.Save(ctx, tx, category)

	//return new category
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	//validate first
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	//set database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//find category by id
	category, err := service.repository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	//update category to database
	category.Name = request.Name
	service.repository.Update(ctx, tx, category)

	//return category update
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	//set database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//find category by id
	category, err := service.repository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	//	delete category
	service.repository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	//set database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// get find by id
	category, err := service.repository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	// set database transactional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//get all category from database
	categories := service.repository.FindAll(ctx, tx)

	//return data
	return helper.ToCategoryResponses(categories)
}
