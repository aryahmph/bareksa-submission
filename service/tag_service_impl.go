package service

import (
	"bareksa-aryayunanta/exception"
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
	if isExist := t.TagRepository.IsExistByName(ctx, tx, tag.Name); !isExist {
		tag = t.TagRepository.Save(ctx, tx, tag)
	} else {
		panic(exception.NewAlreadyExistError("category is already exist"))
	}

	return helper.ToTagResponse(tag)
}

func (t *TagServiceImpl) FindAll(ctx context.Context) []web.TagResponse {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tags := t.TagRepository.FindAll(ctx, tx)

	return helper.ToTagResponses(tags)
}

func (t *TagServiceImpl) FindByName(ctx context.Context, tagName string) web.TagResponse {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tag, err := t.TagRepository.FindByName(ctx, tx, tagName)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTagResponse(tag)
}

func (t *TagServiceImpl) IsExistByName(ctx context.Context, tagName string) bool {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	return t.TagRepository.IsExistByName(ctx, tx, tagName)
}
