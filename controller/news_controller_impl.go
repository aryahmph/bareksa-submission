package controller

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/web"
	"bareksa-aryayunanta/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type NewsControllerImpl struct {
	NewsService service.NewsService
}

func NewNewsControllerImpl(newsService service.NewsService) *NewsControllerImpl {
	return &NewsControllerImpl{NewsService: newsService}
}

func (n *NewsControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	newsResponses := n.NewsService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   newsResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
