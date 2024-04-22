package logic

import (
	"context"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminCheckTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminCheckTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminCheckTaskLogic {
	return &AdminCheckTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminCheckTaskLogic) AdminCheckTask(in *task.AdminCheckTaskRequest) (*task.Empty, error) {
	// todo: add your logic here and delete this line

	return &task.Empty{}, nil
}
