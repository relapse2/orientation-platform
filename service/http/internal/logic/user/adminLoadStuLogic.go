package user

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoadStuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoadStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoadStuLogic {
	return &AdminLoadStuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoadStuLogic) AdminLoadStu(req *types.AdminLoadStuRequest) (resp *types.AdminLoadStuReply, err error) {
	// todo: add your logic here and delete this line

	return
}
