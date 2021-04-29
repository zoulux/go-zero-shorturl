package logic

import (
	"context"
	"shorturl/rpc/transform/transformclient"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShortenLogic {
	return ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req types.ShortenReq) (*types.ShortenResp, error) {
	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformclient.ShortenReq{Url: req.Url})
	if err != nil {
		return &types.ShortenResp{}, err
	}
	return &types.ShortenResp{Shorten: resp.Shorten}, nil
}
