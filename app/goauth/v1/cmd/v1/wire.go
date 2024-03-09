//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"goauth/app/goauth/v1/internal/biz"
	"goauth/app/goauth/v1/internal/conf"
	"goauth/app/goauth/v1/internal/server"
	"goauth/app/goauth/v1/internal/service"
	"goauth/app/goauth/v1/middleware/registry"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		registry.ProviderSet,
		newApp))
}
