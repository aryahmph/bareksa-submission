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

func (t *TopicRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, topicName string) (domain.Topic, error) {
	panic("implement me")
}

func (t *TopicRepositoryImpl) IsExistByName(ctx context.Context, tx *sql.Tx, topicName string) bool {
	SQL := "SELECT id FROM topics WHERE name = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, topicName)
	helper.PanicIfError(err)
	defer rows.Close()

	return rows.Next()
}
