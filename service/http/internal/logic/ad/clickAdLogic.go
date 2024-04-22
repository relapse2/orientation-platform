package ad

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClickAdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClickAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickAdLogic {
	return &ClickAdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClickAdLogic) ClickAd(req *types.ClickAdRequest) (resp *types.ClickAdReply, err error) {
	// todo: add your logic here and delete this line

	return
}
