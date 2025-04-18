//go:build wireinject
// +build wireinject

package main

import (
	"deployment-test/controller"
	"deployment-test/helper"
	"deployment-test/repository"
	"deployment-test/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var ServerSet = wire.NewSet(
	repository.NewRepositoryImpl,
	service.NewService,
	controller.NewController,
	NewServer,
	wire.Bind(new(http.Handler), new(*httprouter.Router)),
	NewRouter,
	helper.NewDb,
)

func IntitializedServer() (*http.Server, func(), error) {
	wire.Build(ServerSet)
	return nil, nil, nil
}
