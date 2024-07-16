// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/GoSimplicity/LinkMe-monorepo/api/user/v1"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/biz"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/conf"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/data"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/server"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, confService *conf.Service, userClient v1.UserClient, logger log.Logger) (*kratos.App, func(), error) {
	db, err := data.NewDB(confData)
	if err != nil {
		return nil, nil, err
	}
	cmdable := data.NewRedis(confData)
	client := data.NewMongoDB(confData)
	dataData, cleanup, err := data.NewData(confData, confService, db, cmdable, logger, client)
	if err != nil {
		return nil, nil, err
	}
	zapLogger := data.NewLogger()
	postData := data.NewPostData(dataData, zapLogger)
	postBiz := biz.NewPostBiz(postData)
	postService := service.NewPostService(postBiz, userClient)
	grpcServer := server.NewGRPCServer(confServer, postService)
	httpServer := server.NewHTTPServer(confServer, postService)
	app := newApp(confService, logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
