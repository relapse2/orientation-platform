package user

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpgradeCharLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpgradeCharLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpgradeCharLogic {
	return &UpgradeCharLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpgradeCharLogic) UpgradeChar(req *types.UpgradeCharRequest) (resp *types.UpgradeCharReply, err error) {
	// todo: add your logic here and delete this line

	return
}
