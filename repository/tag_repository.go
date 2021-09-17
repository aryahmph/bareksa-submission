package repository

import (
	"bareksa-aryayunanta/model/domain"
	"context"
	"database/sql"
)

type TagRepository interface {
	Save(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag
	Update(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag
	Delete(ctx context.Context, tx *sql.Tx, tag domain.Tag)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Tag
	FindById(ctx context.Context, tx *sql.Tx, tagId uint32) domain.Tag
	IsExistByName(ctx context.Context, tx *sql.Tx, tagName string) bool
}
