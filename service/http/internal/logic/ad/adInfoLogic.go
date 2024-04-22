package ad

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdInfoLogic {
	return &AdInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdInfoLogic) AdInfo(req *types.AdInfoRequest) (resp *types.AdInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
