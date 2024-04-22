package logic

import (
	"context"

	"orientation-platform/service/rpc/ad/internal/svc"
	"orientation-platform/service/rpc/ad/types/ad"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClickAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClickAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickAdLogic {
	return &ClickAdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClickAdLogic) ClickAd(in *ad.ClickAdRequest) (*ad.Empty, error) {
	// todo: add your logic here and delete this line

	return &ad.Empty{}, nil
}
