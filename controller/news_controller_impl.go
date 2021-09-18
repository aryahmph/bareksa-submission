package controller

import (
	"bareksa-aryayunanta/exception"
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/domain"
	"bareksa-aryayunanta/model/web"
	"bareksa-aryayunanta/service"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"strings"
)

type NewsControllerImpl struct {
	NewsService  service.NewsService
	TopicService service.TopicService
	TagService   service.TagService
}

func NewNewsControllerImpl(newsService service.NewsService, topicService service.TopicService, tagService service.TagService) *NewsControllerImpl {
	return &NewsControllerImpl{NewsService: newsService, TopicService: topicService, TagService: tagService}
}

func (n *NewsControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	news := domain.News{}

	news.Title = strings.Trim(request.PostFormValue("title"), " ")
	news.TopicName = strings.Trim(request.PostFormValue("topic_name"), " ")
	news.ShortDesc = strings.Trim(request.PostFormValue("short_desc"), " ")
	news.Content = strings.Trim(request.PostFormValue("content"), " ")
	news.Writer = strings.Trim(request.PostFormValue("writer"), " ")
	news.Tags = strings.Trim(request.PostFormValue("tags"), " ")

	// Check topic exist or not
	if isTopicExist := n.TopicService.IsExistByName(ctx, news.TopicName); !isTopicExist {
		panic(exception.NewNotFoundError("topic is not found"))
	}

	// Move to uploads directory
	// Change image name
	file, fileHeader, err := request.FormFile("image")
	helper.PanicIfError(err)
	fileName := helper.RandomString(fileHeader.Filename, len(news.Title))
	fileDestination, err := os.Create("./uploads/" + fileName)
	helper.PanicIfError(err)
	_, err = io.Copy(fileDestination, file)
	helper.PanicIfError(err)

	// Check tag
	tags := strings.Split(news.Tags, ",")
	for _, tag := range tags {
		tag = strings.Trim(tag, " ")
		if isTagExist := n.TagService.IsExistByName(ctx, tag); !isTagExist {
			n.TagService.Create(ctx, web.TagCreateRequest{Name: tag})
		}
	}

	// Create news
	newsResponse := n.NewsService.Create(ctx, news)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   newsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
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
