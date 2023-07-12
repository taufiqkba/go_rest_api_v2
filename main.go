package main

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/taufiqkba/go_rest_api_v2/app"
	"github.com/taufiqkba/go_rest_api_v2/controller"
	"github.com/taufiqkba/go_rest_api_v2/repository"
	"github.com/taufiqkba/go_rest_api_v2/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}
