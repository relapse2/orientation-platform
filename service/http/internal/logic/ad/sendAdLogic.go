package ad

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendAdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendAdLogic {
	return &SendAdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendAdLogic) SendAd(req *types.SendAdRequest) (resp *types.SendAdReply, err error) {
	// todo: add your logic here and delete this line

	return
}
