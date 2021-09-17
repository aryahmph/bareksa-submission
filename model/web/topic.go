package web

type TopicCreateRequest struct {
	Name string `validate:"required,min=2,max=50"`
}

type TopicResponse struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}
