package task

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDoTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoTaskLogic {
	return &DoTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DoTaskLogic) DoTask(req *types.DoTaskRequest) (resp *types.DoTaskReply, err error) {
	// todo: add your logic here and delete this line

	return
}
