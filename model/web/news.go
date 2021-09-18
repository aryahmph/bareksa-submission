package web

type NewsCreateRequest struct {
	Title       string `validate:"required,min=1,max=255" json:"title"`
	Description string `validate:"required,min=1,max=255" json:"description"`
	Content     string `validate:"required,min=1,max=65535" json:"content"`
	Topic       string `validate:"required,min=1,max=255" json:"topic"`
	Writer      string `validate:"required,min=1,max=255" json:"writer"`
	Tags        string `validate:"required,min=1,max=255" json:"tags"`
}

type NewsResponse struct {
	Id          uint32   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Topic       string   `json:"topic"`
	Writer      string   `json:"writer"`
	Tags        []string `json:"tags"`
	ImageURL    string   `json:"image_url"`
}

type ListNewsResponses struct {
	Id          uint32 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	ImageURL    string `json:"image_url"`
}
