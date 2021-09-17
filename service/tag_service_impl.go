package service

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/domain"
	"bareksa-aryayunanta/model/web"
	"bareksa-aryayunanta/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type TagServiceImpl struct {
	TagRepository repository.TagRepository
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewTagServiceImpl(tagRepository repository.TagRepository, DB *sql.DB, validate *validator.Validate) *TagServiceImpl {
	return &TagServiceImpl{TagRepository: tagRepository, DB: DB, Validate: validate}
}

func (t *TagServiceImpl) Create(ctx context.Context, request web.TagCreateRequest) web.TagResponse {
	err := t.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tag := domain.Tag{Name: request.Name}
	tag = t.TagRepository.Save(ctx, tx, tag)

	return helper.ToTagResponse(tag)
}
