package service

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/web"
	"bareksa-aryayunanta/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type NewsServiceImpl struct {
	NewsRepository repository.NewsRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewNewsServiceImpl(newsRepository repository.NewsRepository, DB *sql.DB, validate *validator.Validate) *NewsServiceImpl {
	return &NewsServiceImpl{NewsRepository: newsRepository, DB: DB, Validate: validate}
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
			Date:        item.PublishedAt.Format("Monday, 02 January 2006"),
			ImageURL:    "http://localhost:8080/uploads/news/" + item.ImageURL,
		}
		responses = append(responses, response)
	}

	return responses
}
