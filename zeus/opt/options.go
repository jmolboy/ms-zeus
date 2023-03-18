package opt

import (
	"time"
)

type Options struct {
	AuthQueryStringName string
	CookieName          string
	CookieExpire        time.Duration
	HomeUrl             string
	IsOpenAuthTest      bool
}

type ServiceConf struct {
	EndPoint  string
	AppKey    string
	AppSecret string
}

type RedisConf struct {
	Addr         string
	Pass         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type HandlerOption func(opt *Options)

func WithAuthQueryStringName(tktReqName string) HandlerOption {
	return func(opt *Options) {
		opt.AuthQueryStringName = tktReqName
	}
}

func WithHomeUrl(home string) HandlerOption {
	return func(opt *Options) {
		opt.HomeUrl = home
	}
}

func WithAuthTest(isOpen bool) HandlerOption {
	return func(opt *Options) {
		opt.IsOpenAuthTest = isOpen
	}
}

func WithCookie(cookieName string, expire time.Duration) HandlerOption {
	return func(opt *Options) {
		opt.CookieName = cookieName
		opt.CookieExpire = expire
	}
}
