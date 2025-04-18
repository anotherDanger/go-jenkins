package service

import (
	"context"
	"deployment-test/web"
)

type Service interface {
	Create(ctx context.Context, request *web.Request) (*web.Response, error)
	FindAll(ctx context.Context) ([]*web.Response, error)
	FindById(ctx context.Context, id int) (*web.Response, error)
	Update(ctx context.Context, request *web.Request) (*web.Response, error)
	Delete(ctx context.Context, id int) error
}
