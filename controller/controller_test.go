package controller

import (
	"bytes"
	"context"
	"deployment-test/web"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (mock *MockService) Create(ctx context.Context, request *web.Request) (*web.Response, error) {
	args := mock.Called(ctx, request)
	return args.Get(0).(*web.Response), nil
}

func (mock *MockService) FindAll(ctx context.Context) ([]*web.Response, error) {
	args := mock.Called(ctx)
	return args.Get(0).([]*web.Response), nil
}

func (mock *MockService) FindById(ctx context.Context, id int) (*web.Response, error) {
	args := mock.Called(ctx, id)
	return args.Get(0).(*web.Response), nil
}

func (mock *MockService) Update(ctx context.Context, request *web.Request) (*web.Response, error) {
	args := mock.Called(ctx, request)
	return args.Get(0).(*web.Response), nil
}

func (mock *MockService) Delete(ctx context.Context, id int) error {
	args := mock.Called(ctx, id)
	return args.Error(0)
}

func TestControllerCreate(t *testing.T) {
	svc := new(MockService)
	ctrl := NewController(svc)

	reqBody := &web.Request{
		Author: "Test",
		Title:  "Sample",
	}

	expectedResponse := &web.Response{
		Id:     1,
		Author: "Test",
		Title:  "Sample",
	}

	svc.On("Create", mock.Anything, reqBody).Return(expectedResponse, nil)

	byteBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/v1/book", bytes.NewReader(byteBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	ctrl.Create(rec, req, httprouter.Params{})

	var webResponse web.WebResponse[*web.Response]
	json.NewDecoder(rec.Body).Decode(&webResponse)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, expectedResponse, webResponse.Data)
}

func TestControllerFindAll(t *testing.T) {
	svc := new(MockService)
	ctrl := NewController(svc)

	expectedResponse := []*web.Response{
		{Id: 1, Author: "Test 1", Title: "Book 1"},
		{Id: 2, Author: "Test 2", Title: "Book 2"},
	}

	svc.On("FindAll", mock.Anything).Return(expectedResponse, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/book", nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	ctrl.FindAll(rec, req, httprouter.Params{})

	var webResponse web.WebResponse[[]*web.Response]
	json.NewDecoder(rec.Body).Decode(&webResponse)

	assert.Equal(t, expectedResponse, webResponse.Data)
}

func TestControllerFIndById(t *testing.T) {
	svc := new(MockService)
	ctrl := NewController(svc)

	expectedResponse := &web.Response{
		Id:     1,
		Author: "Test",
		Title:  "Book",
	}

	svc.On("FindById", mock.Anything, 1).Return(expectedResponse, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/book/:id", nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	ctrl.FindById(rec, req, httprouter.Params{httprouter.Param{
		Key:   "id",
		Value: "1",
	}})

	var webResponse web.WebResponse[*web.Response]
	json.NewDecoder(rec.Body).Decode(&webResponse)

	assert.Equal(t, expectedResponse, webResponse.Data)
}
