package repository

import (
	"context"
	"deployment-test/domain"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	repo := &RepositoryImpl{}
	entity := &domain.Domain{
		Id:     1,
		Author: "Andhika",
		Title:  "Belajar Golang Dasar",
	}
	mock.ExpectBegin()
	mock.ExpectExec("insert into books").WithArgs(entity.Author, entity.Title).WillReturnResult(sqlmock.NewResult(1, 1))

	tx, err := db.Begin()
	if err != nil {
		t.Fatal("cannot start tx")
	}
	defer tx.Rollback()

	result, err := repo.Create(context.Background(), tx, entity)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Andhika", result.Author)
	assert.Equal(t, "Belajar Golang Dasar", result.Title)

	mock.ExpectationsWereMet()
}

func TestRepositoryFindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	repo := &RepositoryImpl{}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("select * from books")).WillReturnRows(sqlmock.NewRows([]string{"id", "author", "title"}).
		AddRow(1, "Andhika", "Belajar PHP Dasar").
		AddRow(2, "Danger", "Belajar Docker Compose"))

	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		t.Fatal("cannot start tx")
	}

	results, err := repo.FindAll(context.Background(), tx)

	assert.NoError(t, err)
	assert.Equal(t, 1, results[0].Id)
	assert.Equal(t, "Andhika", results[0].Author)
	assert.Equal(t, "Belajar PHP Dasar", results[0].Title)

	assert.Equal(t, 2, results[1].Id)
	assert.Equal(t, "Danger", results[1].Author)
	assert.Equal(t, "Belajar Docker Compose", results[1].Title)

	mock.ExpectationsWereMet()
}

func TestRepositoryFindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	repo := &RepositoryImpl{}
	domain := &domain.Domain{}
	domain.Id = 1

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("select * from books where id")).WithArgs(domain.Id).WillReturnRows(sqlmock.NewRows([]string{"Id", "Author", "Title"}).
		AddRow(1, "Andhika", "Book 1"))

	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
	result, err := repo.FindById(context.Background(), tx, domain.Id)

	assert.NoError(t, err)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Andhika", result.Author)
	assert.Equal(t, "Book 1", result.Title)

}

func TestRepositoryUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	repo := &RepositoryImpl{}
	domain := &domain.Domain{
		Author: "Danger",
		Title:  "Sample 1",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("update books set author = ?, title = ? where id")).WithArgs(domain.Author, domain.Title, domain.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	defer tx.Rollback()

	result, err := repo.Update(context.Background(), tx, domain)

	assert.NoError(t, err)
	assert.Equal(t, 0, result.Id)
	assert.Equal(t, "Danger", result.Author)
	assert.Equal(t, "Sample 1", result.Title)

	mock.ExpectationsWereMet()
}

func TestRepositoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("delete from books where id")).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := &RepositoryImpl{}

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	err = repo.Delete(context.Background(), tx, 1)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	mock.ExpectationsWereMet()
}
