package domain

import "time"

type News struct {
	ID                     uint32
	Title                  string `validate:"required,min=1,max=255"`
	ShortDesc              string `validate:"required,min=1,max=255"`
	Content                string `validate:"required,min=1,max=65535"`
	TopicName              string `validate:"required,min=1,max=255"`
	Writer                 string `validate:"required,min=1,max=255"`
	ImageURL, Status, Tags string
	PublishedAt, CreatedAt time.Time
	UpdatedAt, DeletedAt   *time.Time
}
