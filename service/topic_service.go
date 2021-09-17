package service

import (
	"bareksa-aryayunanta/model/web"
	"context"
)

type TopicService interface {
	Create(ctx context.Context, request web.TopicCreateRequest) web.TopicResponse
	FindAll(ctx context.Context) []web.TopicResponse
}
