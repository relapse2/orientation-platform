package logic

import (
	"context"

	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminByNameLogic {
	return &GetAdminByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminByNameLogic) GetAdminByName(in *user.GetAdminByNameRequest) (*user.GetAdminReply, error) {
	// todo: add your logic here and delete this line

	return &user.GetAdminReply{}, nil
}
