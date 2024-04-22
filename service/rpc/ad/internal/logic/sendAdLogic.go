package logic

import (
	"context"

	"orientation-platform/service/rpc/ad/internal/svc"
	"orientation-platform/service/rpc/ad/types/ad"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendAdLogic {
	return &SendAdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendAdLogic) SendAd(in *ad.SendAdRequest) (*ad.SendAdReply, error) {
	// todo: add your logic here and delete this line

	return &ad.SendAdReply{}, nil
}
