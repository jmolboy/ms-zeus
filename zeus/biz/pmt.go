package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type AccessToken struct {
	Id          int64
	Uid         int64
	UserId      string
	UserName    string
	RoleId      int64
	AccessToken string
	AppId       int64
	AppUniqId   string
	AppName     string
	Pmts        map[string]int64
	CallBackUrl string
	LoginUrl    string
	LogoutUrl   string
}

type PmtRepo interface {
	Auth(ctx context.Context) (accToken *AccessToken, err error)
	Pmt(ctx context.Context, keys []string) (accToken *AccessToken, err error)
}

type PmtUseCase struct {
	repo PmtRepo
	log  *log.Helper
}

func NewPmtUseCase(repo PmtRepo, logger log.Logger) *PmtUseCase {
	return &PmtUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *PmtUseCase) Auth(ctx context.Context) (accToken *AccessToken, err error) {
	return uc.repo.Auth(ctx)
}

func (uc *PmtUseCase) Pmt(ctx context.Context, keys ...string) (accToken *AccessToken, err error) {
	return uc.repo.Pmt(ctx, keys)
}
