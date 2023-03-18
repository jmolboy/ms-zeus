//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package zeus

import (
	"github.com/jmolboy/ms-zeus/zeus/biz"
	"github.com/jmolboy/ms-zeus/zeus/data"
	"github.com/jmolboy/ms-zeus/zeus/opt"
	"github.com/jmolboy/ms-zeus/zeus/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*opt.ServiceConf, *opt.RedisConf, log.Logger) (*ZeusApp, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, newZeusApp))
}
