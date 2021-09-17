package service

import (
	"bareksa-aryayunanta/model/web"
	"context"
)

type TagService interface {
	Create(ctx context.Context, request web.TagCreateRequest) web.TagResponse
}
