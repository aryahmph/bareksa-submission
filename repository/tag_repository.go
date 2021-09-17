package repository

import (
	"bareksa-aryayunanta/model/domain"
	"context"
	"database/sql"
)

type TagRepository interface {
	Save(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Tag
	FindByName(ctx context.Context, tx *sql.Tx, tagName string) (domain.Tag, error)
	IsExistByName(ctx context.Context, tx *sql.Tx, tagName string) bool
}
