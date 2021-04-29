package logic

import (
	"context"
	"shorturl/rpc/transform/transformclient"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExpandLogic {
	return ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req types.ExpandReq) (*types.ExpandResp, error) {
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformclient.ExpandReq{
		Shorten: req.Shorten,
	})

	if err != nil {
		return &types.ExpandResp{}, err
	}

	return &types.ExpandResp{
		Url: resp.Url,
	}, nil
}
