//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	userpb "github.com/GoSimplicity/LinkMe-monorepo/api/user/v1"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/biz"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/conf"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/data"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/server"
	"github.com/GoSimplicity/LinkMe/app/linkme-post/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Service, userpb.UserClient, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
