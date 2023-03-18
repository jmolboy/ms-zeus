package data

import (
	"context"
	"fmt"
	"strconv"

	xerrors "github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmolboy/ms-zeus/zeus/biz"
)

const usr_pmt_key = "zeus:user:%dpmt:%d"

var _ biz.UsrPmtRepo = (*usrPmtRepo)(nil)

type usrPmtRepo struct {
	data *Data
	log  *log.Helper
}

func NewUsrPmtRepo(data *Data, logger log.Logger) biz.UsrPmtRepo {
	return &usrPmtRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/app")),
	}
}

func (repo *usrPmtRepo) usrPmtKey(uid int64, appId int64) string {
	return fmt.Sprintf(usr_pmt_key, uid, appId)
}

func (repo *usrPmtRepo) Save(ctx context.Context, uid, appId int64, pmts map[string]int64) (err error) {
	// 存入redis
	key := repo.usrPmtKey(uid, appId)
	data := mapToRdsPair(pmts)
	err = repo.data.rdb.HMSet(ctx, key, data...).Err()
	if err != nil {
		err = xerrors.Wrapf(err, "redis存储失败uid:%d appId:%d", uid, appId)
	}
	return
}

func (repo *usrPmtRepo) Pmt(ctx context.Context, uid int64, appId int64, keys ...string) (res map[string]int64, err error) {
	key := repo.usrPmtKey(uid, appId)

	data, err := repo.data.rdb.HMGet(ctx, key, keys...).Result()
	if err != nil {
		err = xerrors.Wrapf(err, "redis鉴权失败uid:%d appId:%d", uid, appId)
		return
	}

	res = map[string]int64{}
	for idx, v := range keys {
		tmpVal := data[idx]
		if tmpVal == nil {
			res[v] = 0
			continue
		}
		val, ok := int64Val(tmpVal)
		if ok == nil {
			res[v] = int64(val)
		} else {
			res[v] = 0
		}
	}
	return
}

func mapToRdsPair(m map[string]int64) []interface{} {
	ret := []interface{}{}
	for k, v := range m {
		ret = append(ret, k, v)
	}
	return ret
}

func int64Val(data interface{}) (int64, error) {
	res := data.(string)
	return strconv.ParseInt(res, 10, 64)
}
