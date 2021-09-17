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
	panic("implement me")
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
