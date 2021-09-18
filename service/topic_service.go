package service

import (
	"bareksa-aryayunanta/model/web"
	"context"
)

type TopicService interface {
	Create(ctx context.Context, request web.TopicCreateRequest) web.TopicResponse
	FindAll(ctx context.Context) []web.TopicResponse
	FindByName(ctx context.Context, topicName string) []web.ListNewsResponses
	IsExistByName(ctx context.Context, topicName string) bool
}
