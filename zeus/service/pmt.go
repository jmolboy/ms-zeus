package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmolboy/ms-zeus/zeus/biz"
)

// Service is service
type PmtService struct {
	pmtUc    *biz.PmtUseCase
	usrPmtUc *biz.UsrPmtUseCase
	log      *log.Helper
}

// New service
func NewPmtService(uc *biz.PmtUseCase, usrPmtUc *biz.UsrPmtUseCase, logger log.Logger) *PmtService {
	return &PmtService{
		pmtUc:    uc,
		usrPmtUc: usrPmtUc,
		log:      log.NewHelper(log.With(logger, "module", "service/pmt")),
	}
}

func (s *PmtService) AccessToken(ctx context.Context) (accessToken *biz.AccessToken, err error) {
	// 调用adm的client方法
	return s.pmtUc.Auth(ctx)
}

func (s *PmtService) Save(ctx context.Context, uid, appId int64, pmts map[string]int64) (err error) {
	return s.usrPmtUc.Save(ctx, uid, appId, pmts)
}

func (s *PmtService) Pmt(ctx context.Context, keys ...string) (res map[string]int64, err error) {
	// 上下文获取用户
	// usr, err := pmtjwt.FromAuthContext(ctx)
	// if err != nil {
	// 	return
	// }

	// if usr.AppId == 0 {
	// 	err = errors.BadRequest("APP_ID_EMPTY", "appid为空，请先完成授权")
	// 	return
	// }
	// usr := pmtjwt.JwtUser{
	// 	Uid:   2,
	// 	AppId: 1,
	// }

	_, err = s.pmtUc.Pmt(ctx, keys...)
	return
}

func (s *PmtService) PmtUsr(ctx context.Context, uid, appId int64, keys ...string) (res map[string]int64, err error) {
	return s.usrPmtUc.Pmt(ctx, uid, appId, keys...)
}
