package app

import (
	"bareksa-aryayunanta/controller"
	"bareksa-aryayunanta/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(tagController controller.TagController, topicController controller.TopicController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/tags", tagController.FindAll)
	router.GET("/api/tags/:tagName", tagController.FindByName)
	router.POST("/api/tags", tagController.Create)

	router.GET("/api/topics", topicController.FindAll)
	router.POST("/api/topics", topicController.Create)

	router.GET("/api/news", topicController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
