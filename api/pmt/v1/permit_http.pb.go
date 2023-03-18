// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.9
// source: pmt/v1/permit.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPermitAuth = "/api.pmt.v1.Permit/Auth"
const OperationPermitPmt = "/api.pmt.v1.Permit/Pmt"
const OperationPermitPmtAll = "/api.pmt.v1.Permit/PmtAll"

type PermitHTTPServer interface {
	// Auth 授权码获取AccessToken
	Auth(context.Context, *AuthRequest) (*AuthReply, error)
	// Pmt 具体权限点验证授权
	Pmt(context.Context, *PmtRequest) (*PmtReply, error)
	// PmtAll 所有权限点的授权结果
	PmtAll(context.Context, *PmtAllRequest) (*PmtAllReply, error)
}

func RegisterPermitHTTPServer(s *http.Server, srv PermitHTTPServer) {
	r := s.Route("/")
	r.POST("/api/permit/v1/auth", _Permit_Auth0_HTTP_Handler(srv))
	r.POST("/api/permit/v1/pmt", _Permit_Pmt0_HTTP_Handler(srv))
	r.POST("/api/permit/v1/pmtall", _Permit_PmtAll0_HTTP_Handler(srv))
}

func _Permit_Auth0_HTTP_Handler(srv PermitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPermitAuth)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Auth(ctx, req.(*AuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthReply)
		return ctx.Result(200, reply)
	}
}

func _Permit_Pmt0_HTTP_Handler(srv PermitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PmtRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPermitPmt)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Pmt(ctx, req.(*PmtRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PmtReply)
		return ctx.Result(200, reply)
	}
}

func _Permit_PmtAll0_HTTP_Handler(srv PermitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PmtAllRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPermitPmtAll)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PmtAll(ctx, req.(*PmtAllRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PmtAllReply)
		return ctx.Result(200, reply)
	}
}

type PermitHTTPClient interface {
	Auth(ctx context.Context, req *AuthRequest, opts ...http.CallOption) (rsp *AuthReply, err error)
	Pmt(ctx context.Context, req *PmtRequest, opts ...http.CallOption) (rsp *PmtReply, err error)
	PmtAll(ctx context.Context, req *PmtAllRequest, opts ...http.CallOption) (rsp *PmtAllReply, err error)
}

type PermitHTTPClientImpl struct {
	cc *http.Client
}

func NewPermitHTTPClient(client *http.Client) PermitHTTPClient {
	return &PermitHTTPClientImpl{client}
}

func (c *PermitHTTPClientImpl) Auth(ctx context.Context, in *AuthRequest, opts ...http.CallOption) (*AuthReply, error) {
	var out AuthReply
	pattern := "/api/permit/v1/auth"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPermitAuth))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermitHTTPClientImpl) Pmt(ctx context.Context, in *PmtRequest, opts ...http.CallOption) (*PmtReply, error) {
	var out PmtReply
	pattern := "/api/permit/v1/pmt"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPermitPmt))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermitHTTPClientImpl) PmtAll(ctx context.Context, in *PmtAllRequest, opts ...http.CallOption) (*PmtAllReply, error) {
	var out PmtAllReply
	pattern := "/api/permit/v1/pmtall"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPermitPmtAll))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}