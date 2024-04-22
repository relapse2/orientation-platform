package task

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminCheckTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminCheckTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminCheckTaskLogic {
	return &AdminCheckTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminCheckTaskLogic) AdminCheckTask(req *types.AdminCheckTaskRequest) (resp *types.AdminCheckTaskReply, err error) {
	// todo: add your logic here and delete this line

	return
}
