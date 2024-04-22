package logic

import (
	"context"

	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCharLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCharLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCharLogic {
	return &UpdateCharLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCharLogic) UpdateChar(in *user.UpgradeCharRequest) (*user.Empty, error) {
	// todo: add your logic here and delete this line

	return &user.Empty{}, nil
}
