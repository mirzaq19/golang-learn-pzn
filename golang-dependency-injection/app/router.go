package app

import (
	"golang-restful-api/controller"
	"golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
