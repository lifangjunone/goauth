package data

import (
	userv1 "goauth/api/user/v1"
	"goauth/app/goauth/v1/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet()

// Data .
type Data struct {
	// TODO wrapped database client
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, uc userv1.UserClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		log: log.NewHelper(logger),
	}, cleanup, nil
}
