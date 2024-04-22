package logic

import (
	"context"

	"orientation-platform/service/rpc/ad/internal/svc"
	"orientation-platform/service/rpc/ad/types/ad"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminSetAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminSetAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminSetAdLogic {
	return &AdminSetAdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminSetAdLogic) AdminSetAd(in *ad.AdminSetAdRequest) (*ad.Empty, error) {
	// todo: add your logic here and delete this line

	return &ad.Empty{}, nil
}
