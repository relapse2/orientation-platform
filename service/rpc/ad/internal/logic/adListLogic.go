package logic

import (
	"context"

	"orientation-platform/service/rpc/ad/internal/svc"
	"orientation-platform/service/rpc/ad/types/ad"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdListLogic {
	return &AdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdListLogic) AdList(in *ad.Empty) (*ad.AdListReply, error) {
	// todo: add your logic here and delete this line

	return &ad.AdListReply{}, nil
}
