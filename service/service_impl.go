package service

import (
	"context"
	"database/sql"
	"deployment-test/domain"
	"deployment-test/helper"
	"deployment-test/repository"
	"deployment-test/web"
)

type ServiceImpl struct {
	Db         *sql.DB
	repository repository.Repository
}

func NewService(Db *sql.DB, repository repository.Repository) Service {
	return &ServiceImpl{
		Db:         Db,
		repository: repository,
	}
}

func (service *ServiceImpl) Create(ctx context.Context, request *web.Request) (*web.Response, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx, &err)

	domainResponse, err := service.repository.Create(ctx, tx, (*domain.Domain)(request))
	if err != nil {
		return nil, err
	}

	response := &web.Response{
		Id:     domainResponse.Id,
		Author: domainResponse.Author,
		Title:  domainResponse.Title,
	}

	return response, nil
}

func (service *ServiceImpl) FindAll(ctx context.Context) ([]*web.Response, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx, &err)

	domainResponses, err := service.repository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	responses := []*web.Response{}

	for _, domain := range domainResponses {
		response := &web.Response{
			Id:     domain.Id,
			Author: domain.Author,
			Title:  domain.Title,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (service *ServiceImpl) FindById(ctx context.Context, id int) (*web.Response, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx, &err)

	domainResponse, err := service.repository.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	response := &web.Response{
		Id:     domainResponse.Id,
		Author: domainResponse.Author,
		Title:  domainResponse.Title,
	}
	return response, nil
}

func (service *ServiceImpl) Update(ctx context.Context, request *web.Request) (*web.Response, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx, &err)

	domainResponse, err := service.repository.Update(ctx, tx, (*domain.Domain)(request))
	if err != nil {
		return nil, err
	}

	response := &web.Response{
		Id:     domainResponse.Id,
		Author: domainResponse.Author,
		Title:  domainResponse.Title,
	}

	return response, nil
}

func (service *ServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.Db.Begin()
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(tx, &err)

	if err := service.repository.Delete(ctx, tx, id); err != nil {
		return err
	}

	return nil
}
