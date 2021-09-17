package repository

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/domain"
	"context"
	"database/sql"
)

type NewsRepositoryImpl struct {
}

func NewNewsRepositoryImpl() *NewsRepositoryImpl {
	return &NewsRepositoryImpl{}
}

func (n *NewsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, news domain.News) domain.News {
	SQL := "INSERT INTO news(id_topic, title, short_desc, content, image_url, writer) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, news.IDTopic, news.Title, news.Content, news.ImageURL, news.Writer)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	news.ID = uint32(id)
	return news
}

func (n *NewsRepositoryImpl) IsExistByID(ctx context.Context, tx *sql.Tx, IDNews uint32) bool {
	SQL := "SELECT id FROM news WHERE id=? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, IDNews)
	helper.PanicIfError(err)
	defer rows.Close()
	return rows.Next()
}
