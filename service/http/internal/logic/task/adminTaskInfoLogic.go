package task

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminTaskInfoLogic {
	return &AdminTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminTaskInfoLogic) AdminTaskInfo(req *types.AdminTaskInfoRequest) (resp *types.AdminTaskInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
