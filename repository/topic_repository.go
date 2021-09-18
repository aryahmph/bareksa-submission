package repository

import (
	"bareksa-aryayunanta/model/domain"
	"context"
	"database/sql"
)

type TopicRepository interface {
	Save(ctx context.Context, tx *sql.Tx, topic domain.Topic) domain.Topic
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Topic
	FindByName(ctx context.Context, tx *sql.Tx, topicName string) []domain.News
	IsExistByName(ctx context.Context, tx *sql.Tx, topicName string) bool
}
