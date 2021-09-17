package controller

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/web"
	"bareksa-aryayunanta/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TopicControllerImpl struct {
	TopicService service.TopicService
}

func NewTopicControllerImpl(topicService service.TopicService) *TopicControllerImpl {
	return &TopicControllerImpl{TopicService: topicService}
}

func (t *TopicControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	topicCreateRequest := web.TopicCreateRequest{}
	helper.ReadFromRequestBody(request, &topicCreateRequest)

	topicResponse := t.TopicService.Create(request.Context(), topicCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   topicResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (t *TopicControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	topicResponses := t.TopicService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   topicResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (t *TopicControllerImpl) FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("implement me")
}
