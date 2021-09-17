package domain

import "time"

type Tag struct {
	ID        uint32
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
