package data

import (
	"context"
	"fmt"

	v1 "github.com/jmolboy/ms-zeus/api/pmt/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmolboy/ms-zeus/zeus/biz"
)

var _ biz.PmtRepo = (*pmtRepo)(nil)

type pmtRepo struct {
	data *Data
	log  *log.Helper
}

func NewPmtRepo(data *Data, logger log.Logger) biz.PmtRepo {
	return &pmtRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/app")),
	}
}

// 获取单个用户
func (r *pmtRepo) Auth(ctx context.Context) (accToken *biz.AccessToken, err error) {
	req := &v1.AuthRequest{}
	reply, err := r.data.pmtCt.Auth(ctx, req)
	if err != nil {
		return
	}

	accToken = &biz.AccessToken{
		Id:          reply.Id,
		Uid:         reply.Uid,
		UserId:      reply.UserId,
		UserName:    reply.UserName,
		RoleId:      reply.RoleId,
		AccessToken: reply.AccessToken,
		AppId:       reply.AppId,
		AppUniqId:   reply.AppUniqId,
		AppName:     reply.AppName,
		Pmts:        reply.Pmts,
		CallBackUrl: reply.CallBackUrl,
		LoginUrl:    reply.LoginUrl,
		LogoutUrl:   reply.LogoutUrl,
	}
	return
}

// 获取单个用户
func (r *pmtRepo) Pmt(ctx context.Context, keys []string) (accToken *biz.AccessToken, err error) {
	req := &v1.PmtRequest{
		Keys: keys,
	}

	reply, err := r.data.pmtCt.Pmt(ctx, req)
	if err != nil {
		return
	}

	fmt.Print(reply.AppId)
	return
}
