package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type UsrPmt struct {
	Id    int64
	Uid   int64
	AppId int64
	Pmts  map[string]int64
}

type UsrPmtRepo interface {
	Save(ctx context.Context, uid, appId int64, pmts map[string]int64) (err error)
	Pmt(ctx context.Context, uid int64, appId int64, keys ...string) (res map[string]int64, err error)
}

type UsrPmtUseCase struct {
	repo UsrPmtRepo
	log  *log.Helper
}

func NewUsrPmtUseCase(repo UsrPmtRepo, logger log.Logger) *UsrPmtUseCase {
	return &UsrPmtUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UsrPmtUseCase) Save(ctx context.Context, uid, appId int64, pmts map[string]int64) (err error) {
	return uc.repo.Save(ctx, uid, appId, pmts)
}

func (uc *UsrPmtUseCase) Pmt(ctx context.Context, uid int64, appId int64, keys ...string) (res map[string]int64, err error) {
	return uc.repo.Pmt(ctx, uid, appId, keys...)
}
