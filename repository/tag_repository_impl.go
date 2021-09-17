package repository

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/domain"
	"context"
	"database/sql"
)

type TagRepositoryImpl struct {
}

func (t *TagRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag {
	SQL := "INSERT INTO tags(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, tag.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	tag.ID = uint32(id)
	return tag
}

func (t *TagRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag {
	panic("implement me")
}

func (t *TagRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, tag domain.Tag) {
	panic("implement me")
}

func (t *TagRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Tag {
	panic("implement me")
}

func (t *TagRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, tagId uint32) domain.Tag {
	panic("implement me")
}