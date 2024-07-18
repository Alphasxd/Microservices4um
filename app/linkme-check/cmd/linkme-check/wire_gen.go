// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/GoSimplicity/LinkMe-microservices/api/post/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/biz"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/conf"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/data"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/server"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, confService *conf.Service, postClient post.PostClient, logger log.Logger) (*kratos.App, func(), error) {
	db, err := data.NewDB(confData)
	if err != nil {
		return nil, nil, err
	}
	zapLogger := data.NewLogger()
	checkData := data.NewCheckData(db, zapLogger)
	checkBiz := biz.NewCheckBiz(checkData)
	checkService := service.NewCheckService(checkBiz, postClient)
	grpcServer := server.NewGRPCServer(confServer, checkService)
	httpServer := server.NewHTTPServer(confServer, checkService)
	app := newApp(confService, logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}
