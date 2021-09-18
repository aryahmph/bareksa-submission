package service

import (
	"bareksa-aryayunanta/model/web"
	"context"
)

type NewsService interface {
	FindAll(ctx context.Context) []web.ListNewsResponses
}
