package logic

import (
	"context"

	"orientation-platform/service/rpc/ad/internal/svc"
	"orientation-platform/service/rpc/ad/types/ad"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckCostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckCostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCostLogic {
	return &CheckCostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckCostLogic) CheckCost(in *ad.CheckCostRequest) (*ad.CheckCostReply, error) {
	// todo: add your logic here and delete this line

	return &ad.CheckCostReply{}, nil
}
