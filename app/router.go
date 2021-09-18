package app

import (
	"bareksa-aryayunanta/controller"
	"bareksa-aryayunanta/exception"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter(tagController controller.TagController, topicController controller.TopicController, newsController controller.NewsController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/tags", tagController.FindAll)
	router.GET("/api/tags/:tagName", tagController.FindByName)
	router.POST("/api/tags", tagController.Create)

	router.GET("/api/topics", topicController.FindAll)
	router.GET("/api/topics/:topicName", topicController.FindByName)
	router.POST("/api/topics", topicController.Create)

	router.POST("/api/news", newsController.Create)
	router.DELETE("/api/news/:newsId", newsController.Delete)
	router.GET("/api/news", newsController.FindAll)
	router.GET("/api/news/:newsId", newsController.FindById)

	directory := http.Dir("./uploads/news/")
	fileServer := http.FileServer(directory)
	router.NotFound = http.StripPrefix("/uploads/news/", fileServer)

	router.PanicHandler = exception.ErrorHandler

	return router
}
