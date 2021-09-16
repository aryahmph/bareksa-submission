package main

import (
	"bareksa-aryayunanta/helper"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}
	err := server.ListenAndServe()
	log.Println("http server started on http://localhost:8080")
	helper.PanicIfError(err)
}
