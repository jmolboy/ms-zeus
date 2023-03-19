package zeus

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	"github.com/jmolboy/ms-adm-pkg/pmtjwt"
	zeusv1 "github.com/jmolboy/ms-zeus/api/zeus/v1"
	"github.com/jmolboy/ms-zeus/zeus/opt"
	"github.com/jmolboy/ms-zeus/zeus/service"
)

const no_service_conf = "NO_PMT_SERVICE_CONF"

const AUTH_PATH = "/api/auth/accesstoken"
const LOG_OUT_PATH = "/api/auth/logout"

type ZeusApp struct {
	pmt    *service.PmtService
	opuser *service.OpUserService
	pmtc   *opt.ServiceConf
	log    *log.Helper
}

func (z *ZeusApp) PmtSvc() *service.PmtService {
	return z.pmt
}

func (z *ZeusApp) OpUserSvc() *service.OpUserService {
	return z.opuser
}

func newZeusApp(pmtc *opt.ServiceConf, pmt *service.PmtService, opuser *service.OpUserService, logger log.Logger) *ZeusApp {
	return &ZeusApp{
		pmt:    pmt,
		opuser: opuser,
		pmtc:   pmtc,
		log:    log.NewHelper(log.With(logger, "module", "pmtapp/app")),
	}
}

func NewZeusApp(pmtc *opt.ServiceConf, rdsc *opt.RedisConf, logger log.Logger) (*ZeusApp, func()) {
	app, cleanup, err := wireApp(pmtc, rdsc, logger)
	if err != nil {
		panic(err)
	}
	return app, cleanup
}

func (z *ZeusApp) Key() (key string, err error) {
	if z.pmtc == nil || z.pmtc.AppKey == "" {
		err = errors.NotFound(no_service_conf, "未配置pmt服务的AppKey")
		return
	}
	key = z.pmtc.AppKey
	return
}

func (z *ZeusApp) Secret() (secret string, err error) {
	if z.pmtc == nil || z.pmtc.AppSecret == "" {
		err = errors.NotFound(no_service_conf, "未配置pmt服务的AppSecret")
		return
	}
	secret = z.pmtc.AppSecret
	return
}

func (z *ZeusApp) AuthHandler(handlerOpts ...opt.HandlerOption) http.Handler {
	opts := &opt.Options{
		CookieName:          "",
		CookieExpire:        1 * time.Hour,
		AuthQueryStringName: "auth",
	}
	for _, o := range handlerOpts {
		o(opts)
	}

	r := mux.NewRouter()
	r.HandleFunc(AUTH_PATH, func(w http.ResponseWriter, r *http.Request) {
		// 从请求获取ticket
		q := r.URL.Query()
		authToken := q.Get(opts.AuthQueryStringName)
		if authToken == "" {
			z.redirect(w, r, opts.HomeUrl)
			return
		}

		// 将authToken传入Context
		ctx := pmtjwt.AuthTokenContext(r.Context(), authToken)
		accessToken, err := z.PmtSvc().AccessToken(ctx)
		if err != nil {
			z.redirect(w, r, opts.HomeUrl)
			return
		}

		// 写到客户端的cookie里去
		if opts.CookieName != "" {
			accCookie := pmtjwt.NewAuthCookie(url.QueryEscape(accessToken.AccessToken), opts.CookieExpire)
			w.Header().Add("Set-Cookie", accCookie.String())
		}
		z.redirect(w, r, accessToken.CallBackUrl)
	}).Methods("GET")

	r.HandleFunc(LOG_OUT_PATH, func(w http.ResponseWriter, r *http.Request) {
		// 写到客户端的cookie里去
		rply, err := z.OpUserSvc().Logout(r.Context(), &zeusv1.LogOutRequest{})
		if err != nil {
			z.writeText(w, err)
			return
		}
		rply.Redirect = opts.HomeUrl

		if opts.CookieName != "" {
			accCookie := pmtjwt.NewAuthCookie(url.QueryEscape(""), time.Second)
			w.Header().Add("Set-Cookie", accCookie.String())
		}
		z.writeText(w, rply)
		return
	}).Methods("GET")

	// 打开测试的cookie，方便app开发调试
	if opts.IsOpenAuthTest {
		r.HandleFunc("/api/auth/test", func(w http.ResponseWriter, r *http.Request) {
			// 写测试的cookie到客户端
			testUsr := pmtjwt.JwtUser{
				Id:       -1,
				UserId:   "test",
				UserName: "测试",
				RoleId:   -1,
				AppId:    -1,
				AppName:  "测试App",
			}

			k, _ := z.Key()
			token, _ := pmtjwt.SignToString(testUsr, k, false)
			if opts.CookieName != "" {
				accCookie := pmtjwt.NewAuthCookie(url.QueryEscape(token), opts.CookieExpire)
				w.Header().Add("Set-Cookie", accCookie.String())
			} else {
				z.writeText(w, errors.Conflict("cookie_name_empty", "CookieName没设置，请使用opt.WithCookie设置"))
			}
		}).Methods("GET")
	}

	return r
}

func (z *ZeusApp) RegisterLoginUserHttpServer(s *khttp.Server) {
	// 还需要一个service
	zeusv1.RegisterOpUserHTTPServer(s, z.opuser)
	// v1.RegisterOpUserHTTPServer(s, service.NewOpUserService())
}

func (z *ZeusApp) redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, 301)
	return
}

func (z *ZeusApp) writeText(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}
