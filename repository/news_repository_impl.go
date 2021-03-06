package repository

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/domain"
	"context"
	"database/sql"
	"errors"
)

type NewsRepositoryImpl struct {
}

func NewNewsRepositoryImpl() *NewsRepositoryImpl {
	return &NewsRepositoryImpl{}
}

func (n *NewsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, news domain.News) domain.News {
	SQL := "INSERT INTO news(topic_name, title, short_desc, content, image_url, writer, status, published_at)\nVALUES (?, ?, ?, ?, ?, ?, 'publish', CURRENT_DATE);"
	result, err := tx.ExecContext(ctx, SQL, news.TopicName, news.Title, news.ShortDesc, news.Content, news.ImageURL, news.Writer)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	news.ID = uint32(id)
	return news
}

func (n *NewsRepositoryImpl) SaveTag(ctx context.Context, tx *sql.Tx, newsId uint32, tagName string) {
	SQL := "INSERT INTO news_tags(id_news, name_tag)\nVALUES (?, ?)"
	_, err := tx.ExecContext(ctx, SQL, newsId, tagName)
	helper.PanicIfError(err)
}

func (n *NewsRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, news domain.News) {
	SQL := "DELETE FROM news_tags WHERE id_news = ?"
	_, err := tx.ExecContext(ctx, SQL, news.ID)
	helper.PanicIfError(err)

	SQL = "DELETE FROM news WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, news.ID)
	helper.PanicIfError(err)
}

func (n *NewsRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.News {
	SQL := "SELECT id, title, short_desc, image_url, published_at\nFROM news\nWHERE status = 'publish'\nORDER BY published_at DESC"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var news []domain.News
	for rows.Next() {
		new2 := domain.News{}
		err := rows.Scan(&new2.ID, &new2.Title, &new2.ShortDesc, &new2.ImageURL, &new2.PublishedAt)
		helper.PanicIfError(err)
		news = append(news, new2)
	}
	return news
}

func (n *NewsRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, newsId uint32) (domain.News, error) {
	SQL := "SELECT title, content, image_url, writer, published_at\nFROM news\nWHERE status = 'publish' AND id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, newsId)
	helper.PanicIfError(err)
	defer rows.Close()

	news := domain.News{}
	if rows.Next() {
		err := rows.Scan(&news.Title, &news.Content, &news.ImageURL, &news.Writer, &news.PublishedAt)
		news.ID = newsId
		helper.PanicIfError(err)
		return news, nil
	} else {
		return news, errors.New("news is not found")
	}
}

func (n *NewsRepositoryImpl) IsExistByID(ctx context.Context, tx *sql.Tx, IDNews uint32) bool {
	SQL := "SELECT id FROM news WHERE id=? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, IDNews)
	helper.PanicIfError(err)
	defer rows.Close()
	return rows.Next()
}
