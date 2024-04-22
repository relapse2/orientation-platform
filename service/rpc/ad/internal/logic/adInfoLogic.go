package logic

import (
	"context"

	"orientation-platform/service/rpc/ad/internal/svc"
	"orientation-platform/service/rpc/ad/types/ad"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdInfoLogic {
	return &AdInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdInfoLogic) AdInfo(in *ad.AdInfoRequest) (*ad.AdInfoReply, error) {
	// todo: add your logic here and delete this line

	return &ad.AdInfoReply{}, nil
}
