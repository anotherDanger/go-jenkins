package service

import (
	"context"
	"database/sql"
	"deployment-test/domain"
	"deployment-test/web"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Create(ctx context.Context, tx *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	args := mock.Called(ctx, tx, entity)
	return args.Get(0).(*domain.Domain), args.Error(1)
}

func (mock *MockRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]*domain.Domain, error) {
	args := mock.Called(ctx, tx)
	return args.Get(0).([]*domain.Domain), args.Error(1)
}

func (mock *MockRepository) FindById(ctx context.Context, tx *sql.Tx, id int) (*domain.Domain, error) {
	args := mock.Called(ctx, tx, id)
	return args.Get(0).(*domain.Domain), args.Error(1)
}

func (mock *MockRepository) Update(ctx context.Context, tx *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	args := mock.Called(ctx, tx, entity)
	return args.Get(0).(*domain.Domain), args.Error(1)
}

func (mock *MockRepository) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	args := mock.Called(ctx, tx, id)
	return args.Error(0)
}

func TestServiceCreate(t *testing.T) {

	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatal("Failed to create mock DB:", err)
	}
	defer db.Close()

	repo := new(MockRepository)

	repo.On("Create", mock.Anything, mock.Anything, mock.AnythingOfType("*domain.Domain")).Return(&domain.Domain{
		Id:     1,
		Author: "Bebas",
		Title:  "Ini juga bebas",
	}, nil)

	svc := NewService(db, repo)

	request := &web.Request{
		Id:     1,
		Author: "Bebas",
		Title:  "Ini juga bebas",
	}
	mockDB.ExpectBegin()

	result, err := svc.Create(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, "Bebas", result.Author)
	assert.Equal(t, "Ini juga bebas", result.Title)

	repo.AssertExpectations(t)
}

func TestServiceFindAll(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	repo := new(MockRepository)
	repo.On("FindAll", mock.Anything, mock.Anything).Return([]*domain.Domain{
		{
			Id:     1,
			Author: "author 1",
			Title:  "book 1",
		},
		{
			Id:     2,
			Author: "author 2",
			Title:  "book 2",
		},
	}, nil)
	mockDB.ExpectBegin()
	svc := NewService(db, repo)

	response, err := svc.FindAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)

	assert.Equal(t, 1, response[0].Id)
	assert.Equal(t, "author 1", response[0].Author)
	assert.Equal(t, "book 1", response[0].Title)

	assert.Equal(t, 2, response[1].Id)
	assert.Equal(t, "author 2", response[1].Author)
	assert.Equal(t, "book 2", response[1].Title)

	repo.AssertExpectations(t)

}

func TestServiceFindById(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	repo := new(MockRepository)
	repo.On("FindById", mock.Anything, mock.Anything, 1).Return(&domain.Domain{
		Id:     1,
		Author: "Test",
		Title:  "judul",
	}, nil)

	mockDB.ExpectBegin()
	svc := NewService(db, repo)
	response, err := svc.FindById(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "Test", response.Author)
	assert.Equal(t, "judul", response.Title)

	repo.AssertExpectations(t)

}

func TestServiceUpdate(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := new(MockRepository)
	repo.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("*domain.Domain")).Return(&domain.Domain{
		Id:     1,
		Author: "update",
		Title:  "test update",
	}, nil)

	mockDB.ExpectBegin()

	request := &web.Request{
		Id:     1,
		Author: "update",
		Title:  "test update",
	}

	svc := NewService(db, repo)
	response, err := svc.Update(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, 1, response.Id)
	assert.Equal(t, "update", response.Author)
	assert.Equal(t, "test update", response.Title)

	repo.AssertExpectations(t)
}

func TestServiceDelete(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()
	repo := new(MockRepository)
	repo.On("Delete", mock.Anything, mock.Anything, 1).Return(nil, nil)

	mockDB.ExpectBegin()
	svc := NewService(db, repo)
	if err := svc.Delete(context.Background(), 1); err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
