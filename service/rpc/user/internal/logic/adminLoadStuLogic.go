package logic

import (
	"context"

	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoadStuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoadStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoadStuLogic {
	return &AdminLoadStuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLoadStuLogic) AdminLoadStu(in *user.AdminLoadStuRequest) (*user.Empty, error) {
	// todo: add your logic here and delete this line

	return &user.Empty{}, nil
}
