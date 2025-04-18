package repository

import (
	"context"
	"database/sql"
	"deployment-test/domain"
)

type RepositoryImpl struct{}

func NewRepositoryImpl() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (repository *RepositoryImpl) Create(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	query := "insert into books(author, title) values(?, ?)"
	result, err := sql.ExecContext(ctx, query, entity.Author, entity.Title)
	if err != nil {
		return nil, err
	}

	lastInsertedId, _ := result.LastInsertId()

	response := &domain.Domain{
		Id:     int(lastInsertedId),
		Author: entity.Author,
		Title:  entity.Title,
	}

	return response, nil
}

func (repository *RepositoryImpl) FindAll(ctx context.Context, sql *sql.Tx) ([]*domain.Domain, error) {
	query := "select * from books"
	rows, err := sql.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []*domain.Domain

	for rows.Next() {
		var bookId int
		var bookAuthor, bookTitle string

		err := rows.Scan(&bookId, &bookAuthor, &bookTitle)
		if err != nil {
			return nil, err
		}

		domain := &domain.Domain{
			Id:     bookId,
			Author: bookAuthor,
			Title:  bookTitle,
		}

		results = append(results, domain)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (repository *RepositoryImpl) FindById(ctx context.Context, sql *sql.Tx, id int) (*domain.Domain, error) {
	query := "select * from books where id = ?"
	row := sql.QueryRowContext(ctx, query, id)

	var bookId int
	var bookAuthor, bookTitle string

	err := row.Scan(&bookId, &bookAuthor, &bookTitle)
	if err != nil {
		return nil, err
	}

	response := &domain.Domain{
		Id:     bookId,
		Author: bookAuthor,
		Title:  bookTitle,
	}

	return response, nil
}

func (repository *RepositoryImpl) Update(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	query := "update books set author = ?, title = ? where id = ?"
	result, err := sql.ExecContext(ctx, query, entity.Author, entity.Title, entity.Id)
	if err != nil {
		return nil, err
	}

	lastInsertedId, _ := result.LastInsertId()

	response := &domain.Domain{
		Id:     int(lastInsertedId),
		Author: entity.Author,
		Title:  entity.Title,
	}

	return response, nil
}

func (repository *RepositoryImpl) Delete(ctx context.Context, sql *sql.Tx, id int) error {
	query := "delete from books where id = ?"
	_, err := sql.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
