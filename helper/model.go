package helper

import (
	"bareksa-aryayunanta/model/domain"
	"bareksa-aryayunanta/model/web"
)

func ToTagResponse(tag domain.Tag) web.TagResponse {
	return web.TagResponse{
		Id:   tag.ID,
		Name: tag.Name,
	}
}

func ToTagResponses(tags []domain.Tag) []web.TagResponse {
	var tagResponses []web.TagResponse
	for _, tag := range tags {
		tagResponses = append(tagResponses, ToTagResponse(tag))
	}
	return tagResponses
}

func ToTopicResponse(topic domain.Topic) web.TopicResponse {
	return web.TopicResponse{
		Id:   topic.Id,
		Name: topic.Name,
	}
}

func ToTopicResponses(topics []domain.Topic) []web.TopicResponse {
	var tagResponses []web.TopicResponse
	for _, topic := range topics {
		tagResponses = append(tagResponses, ToTopicResponse(topic))
	}
	return tagResponses
}
