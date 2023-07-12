package main

import (
	"fmt"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/taufiqkba/go_rest_api_v2/app"
	"github.com/taufiqkba/go_rest_api_v2/controller"
	"github.com/taufiqkba/go_rest_api_v2/helper"
	"github.com/taufiqkba/go_rest_api_v2/repository"
	"github.com/taufiqkba/go_rest_api_v2/service"
	"net/http"
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

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	fmt.Println("Running ...")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
