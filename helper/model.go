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
