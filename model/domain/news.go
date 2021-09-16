package domain

import "time"

type News struct {
	ID, IDTopic                                         uint32
	Title, ShortDesc, Content, ImageURL, Writer, Status string
	CreatedAt                                           time.Time
	UpdatedAt, DeletedAt                                *time.Time
}
