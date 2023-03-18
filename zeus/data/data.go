package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/jmolboy/ms-adm-pkg/pmtjwt"
	api "github.com/jmolboy/ms-zeus/api/pmt/v1"
	"github.com/jmolboy/ms-zeus/zeus/opt"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewRds,
	NewPmtClient,
	NewPmtRepo,
	NewUsrPmtRepo,
)

// Data .
type Data struct {
	log   *log.Helper
	pmtCt api.PermitHTTPClient
	rdb   redis.Cmdable
}

func NewData(pmtCt api.PermitHTTPClient, rdsCmd redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	d := &Data{
		pmtCt: pmtCt,
		rdb:   rdsCmd,
		log:   l,
	}

	cleanup := func() {

	}

	return d, cleanup, nil
}

func NewPmtClient(svcConf *opt.ServiceConf, logger log.Logger) (api.PermitHTTPClient, func()) {
	conn, err := http.NewClient(
		context.Background(),
		http.WithEndpoint(svcConf.EndPoint),
		http.WithMiddleware(
			recovery.Recovery(),
			pmtjwt.PmtAuthClientMiddleware(svcConf.AppKey, svcConf.AppSecret),
		),
	)
	if err != nil {
		panic(err)
	}
	c := api.NewPermitHTTPClient(conn)

	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}

	return c, cleanup
}

func NewRds(conf *opt.RedisConf, logger log.Logger) (redis.Cmdable, func()) {
	log := log.NewHelper(log.With(logger, "module", "zeus/data/rds"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	_, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	cleanup := func() {
		if err := client.Close(); err != nil {
			log.Error(err)
		}
	}
	return client, cleanup
}
