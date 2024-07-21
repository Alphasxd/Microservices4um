// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.27.1
// source: api/interactive/v1/interactive.proto

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

const OperationInteractiveGetInteractive = "/api.interactive.v1.Interactive/GetInteractive"
const OperationInteractiveListInteractive = "/api.interactive.v1.Interactive/ListInteractive"

type InteractiveHTTPServer interface {
	GetInteractive(context.Context, *GetInteractiveRequest) (*GetInteractiveReply, error)
	ListInteractive(context.Context, *ListInteractiveRequest) (*ListInteractiveReply, error)
}

func RegisterInteractiveHTTPServer(s *http.Server, srv InteractiveHTTPServer) {
	r := s.Route("/")
	r.GET("/get/{postId}", _Interactive_GetInteractive0_HTTP_Handler(srv))
	r.POST("/list", _Interactive_ListInteractive0_HTTP_Handler(srv))
}

func _Interactive_GetInteractive0_HTTP_Handler(srv InteractiveHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetInteractiveRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationInteractiveGetInteractive)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetInteractive(ctx, req.(*GetInteractiveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetInteractiveReply)
		return ctx.Result(200, reply)
	}
}

func _Interactive_ListInteractive0_HTTP_Handler(srv InteractiveHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListInteractiveRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationInteractiveListInteractive)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListInteractive(ctx, req.(*ListInteractiveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListInteractiveReply)
		return ctx.Result(200, reply)
	}
}

type InteractiveHTTPClient interface {
	GetInteractive(ctx context.Context, req *GetInteractiveRequest, opts ...http.CallOption) (rsp *GetInteractiveReply, err error)
	ListInteractive(ctx context.Context, req *ListInteractiveRequest, opts ...http.CallOption) (rsp *ListInteractiveReply, err error)
}

type InteractiveHTTPClientImpl struct {
	cc *http.Client
}

func NewInteractiveHTTPClient(client *http.Client) InteractiveHTTPClient {
	return &InteractiveHTTPClientImpl{client}
}

func (c *InteractiveHTTPClientImpl) GetInteractive(ctx context.Context, in *GetInteractiveRequest, opts ...http.CallOption) (*GetInteractiveReply, error) {
	var out GetInteractiveReply
	pattern := "/get/{postId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationInteractiveGetInteractive))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *InteractiveHTTPClientImpl) ListInteractive(ctx context.Context, in *ListInteractiveRequest, opts ...http.CallOption) (*ListInteractiveReply, error) {
	var out ListInteractiveReply
	pattern := "/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationInteractiveListInteractive))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
