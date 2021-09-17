package controller

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/web"
	"bareksa-aryayunanta/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TagControllerImpl struct {
	TagService service.TagService
}

func NewTagControllerImpl(tagService service.TagService) *TagControllerImpl {
	return &TagControllerImpl{TagService: tagService}
}

func (t *TagControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tagCreateRequest := web.TagCreateRequest{}
	helper.ReadFromRequestBody(request, &tagCreateRequest)

	tagResponse := t.TagService.Create(request.Context(), tagCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tagResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (t *TagControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tagResponses := t.TagService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tagResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
