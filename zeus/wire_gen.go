// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package zeus

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmolboy/ms-zeus/zeus/biz"
	"github.com/jmolboy/ms-zeus/zeus/data"
	"github.com/jmolboy/ms-zeus/zeus/opt"
	"github.com/jmolboy/ms-zeus/zeus/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(serviceConf *opt.ServiceConf, redisConf *opt.RedisConf, logger log.Logger) (*ZeusApp, func(), error) {
	permitHTTPClient, cleanup := data.NewPmtClient(serviceConf, logger)
	cmdable, cleanup2 := data.NewRds(redisConf, logger)
	dataData, cleanup3, err := data.NewData(permitHTTPClient, cmdable, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pmtRepo := data.NewPmtRepo(dataData, logger)
	pmtUseCase := biz.NewPmtUseCase(pmtRepo, logger)
	usrPmtRepo := data.NewUsrPmtRepo(dataData, logger)
	usrPmtUseCase := biz.NewUsrPmtUseCase(usrPmtRepo, logger)
	pmtService := service.NewPmtService(pmtUseCase, usrPmtUseCase, logger)
	opUserService := service.NewOpUserService()
	zeusApp := newZeusApp(serviceConf, pmtService, opUserService, logger)
	return zeusApp, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
