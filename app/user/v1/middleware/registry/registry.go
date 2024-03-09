package registry

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

type Registry interface {
	Registry() *registry.Registrar
	Discovery() *registry.Discovery
}

// ProviderSet is registry providers.
var ProviderSet = wire.NewSet(NewConsulRegistry)
