package service

import (
	"bareksa-aryayunanta/exception"
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/domain"
	"bareksa-aryayunanta/model/web"
	"bareksa-aryayunanta/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type TopicServiceImpl struct {
	TopicRepository repository.TopicRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewTopicServiceImpl(topicRepository repository.TopicRepository, DB *sql.DB, validate *validator.Validate) *TopicServiceImpl {
	return &TopicServiceImpl{TopicRepository: topicRepository, DB: DB, Validate: validate}
}

func (t *TopicServiceImpl) Create(ctx context.Context, request web.TopicCreateRequest) web.TopicResponse {
	err := t.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	topic := domain.Topic{Name: request.Name}
	if isExist := t.TopicRepository.IsExistByName(ctx, tx, topic.Name); !isExist {
		topic = t.TopicRepository.Save(ctx, tx, topic)
	} else {
		panic(exception.NewAlreadyExistError("topic is already exist"))
	}

	return helper.ToTopicResponse(topic)
}

func (t *TopicServiceImpl) FindByName(ctx context.Context, topicName string) []web.ListNewsResponses {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	if isNewsExist := t.TopicRepository.IsExistByName(ctx, tx, topicName); !isNewsExist {
		panic(exception.NewNotFoundError("topic is not found"))
	}

	news := t.TopicRepository.FindByName(ctx, tx, topicName)
	var responses []web.ListNewsResponses
	for _, item := range news {
		response := web.ListNewsResponses{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.ShortDesc,
			Date:        helper.TimeFormat(item.PublishedAt),
			ImageURL:    "http://localhost:8080/uploads/news/" + item.ImageURL,
		}
		responses = append(responses, response)
	}

	return responses
}

func (t *TopicServiceImpl) FindAll(ctx context.Context) []web.TopicResponse {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	topics := t.TopicRepository.FindAll(ctx, tx)

	return helper.ToTopicResponses(topics)
}

func (t *TopicServiceImpl) IsExistByName(ctx context.Context, topicName string) bool {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	return t.TopicRepository.IsExistByName(ctx, tx, topicName)
}
