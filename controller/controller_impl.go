package controller

import (
	"deployment-test/service"
	"deployment-test/web"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ControllerImpl struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &ControllerImpl{
		service: service,
	}
}

func (controller *ControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := &web.Request{}
	json.NewDecoder(r.Body).Decode(request)
	response, err := controller.service.Create(r.Context(), request)
	if err != nil {
		webResponse := &web.WebResponse[*web.Response]{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(400),
			Message: err.Error(),
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(webResponse)
		return
	}

	webResponse := &web.WebResponse[*web.Response]{
		Code:    http.StatusCreated,
		Status:  http.StatusText(http.StatusCreated),
		Message: "Created",
		Data:    response,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(webResponse)

}

func (controller *ControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	responses, err := controller.service.FindAll(r.Context())
	if err != nil {
		webResponses := &web.WebResponse[*web.Response]{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Error",
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(webResponses)
		return
	}

	webResponses := &web.WebResponse[[]*web.Response]{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "OK",
		Data:    responses,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponses)
}

func (controller *ControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookId := p.ByName("id")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("cannot convert to int")
		return
	}

	response, err := controller.service.FindById(r.Context(), id)
	if err != nil {
		webResponse := &web.WebResponse[*web.Response]{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Error",
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(webResponse)
		return
	}

	webResponse := &web.WebResponse[*web.Response]{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "OK",
		Data:    response,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *ControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	BookId := p.ByName("id")
	id, err := strconv.Atoi(BookId)
	if err != nil {
		fmt.Println("cannot convert to int")
		return
	}

	webRequest := &web.Request{}
	json.NewDecoder(r.Body).Decode(webRequest)
	webRequest.Id = id

	response, err := controller.service.Update(r.Context(), webRequest)
	if err != nil {
		webResponse := &web.WebResponse[*web.Response]{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Error",
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(webResponse)
		return
	}

	webResponse := &web.WebResponse[*web.Response]{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "OK",
		Data:    response,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *ControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	BookId := p.ByName("id")
	id, err := strconv.Atoi(BookId)
	if err != nil {
		fmt.Println("cannot convert to int")
		return
	}

	err = controller.service.Delete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
