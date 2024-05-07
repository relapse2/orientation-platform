package task

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminSetTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminSetTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminSetTaskLogic {
	return &AdminSetTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminSetTaskLogic) AdminSetTask(req *types.AdminSetTaskRequest) (resp *types.AdminSetTaskReply, err error) {
	// todo: add your logic here and delete this line

	return
}
