// Code generated by goctl. DO NOT EDIT!
// Source: transform.proto

//go:generate mockgen -destination ./transform_mock.go -package transformclient -source $GOFILE

package transformclient

import (
	"context"

	"shorturl/rpc/transform/transform"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	ExpandResp  = transform.ExpandResp
	ShortenReq  = transform.ShortenReq
	ShortenResp = transform.ShortenResp
	Request     = transform.Request
	Response    = transform.Response
	ExpandReq   = transform.ExpandReq

	Transform interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
		Expand(ctx context.Context, in *ExpandReq) (*ExpandResp, error)
		Shorten(ctx context.Context, in *ShortenReq) (*ShortenResp, error)
	}

	defaultTransform struct {
		cli zrpc.Client
	}
)

func NewTransform(cli zrpc.Client) Transform {
	return &defaultTransform{
		cli: cli,
	}
}

func (m *defaultTransform) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := transform.NewTransformClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultTransform) Expand(ctx context.Context, in *ExpandReq) (*ExpandResp, error) {
	client := transform.NewTransformClient(m.cli.Conn())
	return client.Expand(ctx, in)
}

func (m *defaultTransform) Shorten(ctx context.Context, in *ShortenReq) (*ShortenResp, error) {
	client := transform.NewTransformClient(m.cli.Conn())
	return client.Shorten(ctx, in)
}