package main

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/middleware"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	logMiddleware := middleware.NewLogMiddleware(router)

	server := http.Server{
		Handler: logMiddleware,
		Addr:    "localhost:8080",
	}
	err := server.ListenAndServe()
	log.Println("http server started on http://localhost:8080")
	helper.PanicIfError(err)
}
