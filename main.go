package main

import (
	"bareksa-aryayunanta/app"
	"bareksa-aryayunanta/controller"
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/middleware"
	"bareksa-aryayunanta/repository"
	"bareksa-aryayunanta/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	tagRepository := repository.NewTagRepositoryImpl()
	tagService := service.NewTagServiceImpl(tagRepository, db, validate)
	tagController := controller.NewTagControllerImpl(tagService)

	topicRepository := repository.NewTopicRepositoryImpl()
	topicService := service.NewTopicServiceImpl(topicRepository, db, validate)
	topicController := controller.NewTopicControllerImpl(topicService)

	newsRepository := repository.NewNewsRepositoryImpl()
	newsService := service.NewNewsServiceImpl(newsRepository, db, validate)
	newsController := controller.NewNewsControllerImpl(newsService, topicService, tagService)

	router := app.NewRouter(tagController, topicController, newsController)
	logMiddleware := middleware.NewLogMiddleware(router)
	authMiddleware := middleware.NewAuthMiddleware(logMiddleware)

	server := http.Server{
		Handler: authMiddleware,
		Addr:    ":8080",
	}
	fmt.Println("-> http server started on http://" + server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
