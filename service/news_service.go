package service

import (
	"bareksa-aryayunanta/model/domain"
	"bareksa-aryayunanta/model/web"
	"context"
)

type NewsService interface {
	Create(ctx context.Context, request domain.News) web.NewsResponse
	FindAll(ctx context.Context) []web.ListNewsResponses
}
