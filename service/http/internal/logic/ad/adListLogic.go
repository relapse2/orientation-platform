package ad

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdListLogic {
	return &AdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdListLogic) AdList(req *types.AdListRequest) (resp *types.AdListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
