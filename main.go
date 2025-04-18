package main

import (
	"deployment-test/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
}

func NewRouter(controller controller.Controller) *httprouter.Router {
	router := httprouter.New()
	router.POST("/v1/book", controller.Create)
	router.GET("/v1/book", controller.FindAll)
	router.GET("/v1/book/:id", controller.FindById)
	router.PUT("/v1/book/:id", controller.Update)
	router.DELETE("/v1/book/:id", controller.Delete)
	return router
}

func main() {
	server, cleanup, err := IntitializedServer()
	if err != nil {
		panic(err)
	}

	defer cleanup()

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
