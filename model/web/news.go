package web

type NewsResponse struct {
	Id          uint32 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Topic       string `json:"topic"`
	Writer      string `json:"writer"`
	Tags        string `json:"tags"`
	ImageURL    string `json:"image_url"`
}

type ListNewsResponses struct {
	Id          uint32 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	ImageURL    string `json:"image_url"`
}
