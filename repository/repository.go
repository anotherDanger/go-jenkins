package repository

import (
	"context"
	"database/sql"
	"deployment-test/domain"
)

type Repository interface {
	Create(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error)
	FindAll(ctx context.Context, sql *sql.Tx) ([]*domain.Domain, error)
	FindById(ctx context.Context, sql *sql.Tx, id int) (*domain.Domain, error)
	Update(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error)
	Delete(ctx context.Context, sql *sql.Tx, id int) error
}
