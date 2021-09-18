package repository

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/domain"
	"context"
	"database/sql"
)

type TopicRepositoryImpl struct {
}

func NewTopicRepositoryImpl() *TopicRepositoryImpl {
	return &TopicRepositoryImpl{}
}

func (t *TopicRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, topic domain.Topic) domain.Topic {
	SQL := "INSERT INTO topics(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, topic.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	topic.Id = uint32(id)
	return topic
}

func (t *TopicRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Topic {
	SQL := "SELECT id, name FROM topics"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var topics []domain.Topic
	for rows.Next() {
		topic := domain.Topic{}
		err := rows.Scan(&topic.Id, &topic.Name)
		helper.PanicIfError(err)
		topics = append(topics, topic)
	}
	return topics
}

func (t *TopicRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, topicName string) []domain.News {
	SQL := "SELECT n.id, n.title, n.short_desc, n.image_url, n.published_at\nFROM news AS n\nJOIN topics t on t.name = n.topic_name\nWHERE n.topic_name = ?"
	rows, err := tx.QueryContext(ctx, SQL, topicName)
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

func (t *TopicRepositoryImpl) IsExistByName(ctx context.Context, tx *sql.Tx, topicName string) bool {
	SQL := "SELECT id FROM topics WHERE name = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, topicName)
	helper.PanicIfError(err)
	defer rows.Close()

	return rows.Next()
}
