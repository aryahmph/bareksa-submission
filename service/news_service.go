package service

import (
	"bareksa-aryayunanta/model/domain"
	"bareksa-aryayunanta/model/web"
	"context"
)

type NewsService interface {
	Create(ctx context.Context, request domain.News) web.NewsResponse
	Delete(ctx context.Context, newsId uint32)
	FindAll(ctx context.Context) []web.ListNewsResponses
	FindById(ctx context.Context, newsId uint32) web.GetNewsResponse
}
