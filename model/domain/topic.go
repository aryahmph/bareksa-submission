package domain

import "time"

type Topic struct {
	Id        uint32
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
