package service

import (
	"bareksa-aryayunanta/model/web"
	"context"
)

type TagService interface {
	Create(ctx context.Context, request web.TagCreateRequest) web.TagResponse
	FindAll(ctx context.Context) []web.TagResponse
	FindByName(ctx context.Context, tagName string) web.TagResponse
}
