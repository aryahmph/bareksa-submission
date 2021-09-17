package app

import (
	"bareksa-aryayunanta/controller"
	"bareksa-aryayunanta/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(tagController controller.TagController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/tags", tagController.FindAll)
	router.POST("/api/tags", tagController.Create)

	router.PanicHandler = exception.ErrorHandler

	return router
}
