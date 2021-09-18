package repository

import (
	"bareksa-aryayunanta/model/domain"
	"context"
	"database/sql"
)

type NewsRepository interface {
	Save(ctx context.Context, tx *sql.Tx, news domain.News) domain.News
	SaveTag(ctx context.Context, tx *sql.Tx, newsId uint32, tagName string)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.News
	IsExistByID(ctx context.Context, tx *sql.Tx, IDNews uint32) bool
}
