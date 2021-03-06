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
	"strings"
)

type NewsServiceImpl struct {
	NewsRepository repository.NewsRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewNewsServiceImpl(newsRepository repository.NewsRepository, DB *sql.DB, validate *validator.Validate) *NewsServiceImpl {
	return &NewsServiceImpl{NewsRepository: newsRepository, DB: DB, Validate: validate}
}

func (n *NewsServiceImpl) Create(ctx context.Context, request domain.News) web.NewsResponse {
	err := n.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := n.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	news := n.NewsRepository.Save(ctx, tx, request)
	tags := strings.Split(news.Tags, ",")
	for _, tag := range tags {
		tag = strings.Trim(tag, " ")
		n.NewsRepository.SaveTag(ctx, tx, news.ID, tag)
	}

	return web.NewsResponse{
		Id:          news.ID,
		Title:       news.Title,
		Description: news.ShortDesc,
		Content:     news.Content,
		Topic:       news.TopicName,
		Writer:      news.Writer,
		Tags:        news.Tags,
		ImageURL:    news.ImageURL,
	}
}

func (n *NewsServiceImpl) Delete(ctx context.Context, newsId uint32) {
	tx, err := n.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	news, err := n.NewsRepository.FindById(ctx, tx, newsId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	n.NewsRepository.Delete(ctx, tx, news)
}

func (n *NewsServiceImpl) FindAll(ctx context.Context) []web.ListNewsResponses {
	tx, err := n.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	news := n.NewsRepository.FindAll(ctx, tx)
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

func (n *NewsServiceImpl) FindById(ctx context.Context, newsId uint32) web.GetNewsResponse {
	tx, err := n.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	news, err := n.NewsRepository.FindById(ctx, tx, newsId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.GetNewsResponse{
		Id:       news.ID,
		Title:    news.Title,
		Date:     helper.TimeFormat(news.PublishedAt),
		Writer:   news.Writer,
		ImageURL: "http://localhost:8080/uploads/news/" + news.ImageURL,
		Content:  news.Content,
	}
}
