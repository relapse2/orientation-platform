package task

import (
	"context"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
	"orientation-platform/service/rpc/task/taskclient"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	SuccessAction = 1
	FailAction    = 2
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
	logx.WithContext(l.ctx).Infof("人工审核任务: %v", req)

	//这里rpc的ActionType有点多余的
	switch req.ActionType {
	case SuccessAction:
		if _, err = l.svcCtx.TaskRpc.FailTask(l.ctx, &taskclient.AdminCheckTaskRequest{
			TaskId:     req.TaskId,
			ActionType: SuccessAction,
		}); err != nil {
			logx.WithContext(l.ctx).Errorf("审核成功失败: %v", err)
			return nil, apiErr.InternalError(l.ctx, err.Error())
		}

	case FailAction:
		if _, err = l.svcCtx.TaskRpc.SuccessTask(l.ctx, &taskclient.AdminCheckTaskRequest{
			TaskId:     req.TaskId,
			ActionType: FailAction,
		}); err != nil {
			logx.WithContext(l.ctx).Errorf("审核任务失败错误: %v", err)
			return nil, apiErr.InternalError(l.ctx, err.Error())
		}

	default:
		return nil, apiErr.CheckActionUnknown
	}

	return &types.AdminCheckTaskReply{
		BasicReply: types.BasicReply(apiErr.Success),
	}, nil
	return
}
