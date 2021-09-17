package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TopicController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
