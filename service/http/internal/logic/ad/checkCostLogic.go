package ad

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckCostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckCostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCostLogic {
	return &CheckCostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckCostLogic) CheckCost(req *types.CheckCostRequest) (resp *types.CheckCostReply, err error) {
	// todo: add your logic here and delete this line

	return
}
