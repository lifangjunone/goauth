package registry

import (
	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	"goauth/app/user/v1/internal/conf"
)

func NewConsulRegistry(conf *conf.Registry) registry.Registrar {
	// new consul client
	consulConf := api.DefaultConfig()
	consulConf.Address = conf.Address
	client, err := api.NewClient(consulConf)
	if err != nil {
		panic(err)
	}
	// new reg with consul client
	reg := consul.New(client)
	return reg
}
