package web

type TagCreateRequest struct {
	Name string `validate:"required,min=2,max=50"`
}

type TagResponse struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}
