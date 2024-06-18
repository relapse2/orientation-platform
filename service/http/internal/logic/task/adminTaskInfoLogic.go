package task

import (
	"context"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/service/rpc/task/types/task"

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
	logx.WithContext(l.ctx).Infof("admin获取任务信息: %v", req)

	//获取任务信息
	TaskInfoReply, err := l.svcCtx.TaskRpc.TaskInfo(l.ctx, &task.TaskInfoRequest{
		TaskId: req.TaskId,
	})
	if rpcErr.Is(err, rpcErr.TaskNotExist) {
		return nil, apiErr.TaskNotFound
	} else if err != nil {
		logx.WithContext(l.ctx).Errorf("admin获取任务信息失败: %v", err)
		return nil, apiErr.InternalError(l.ctx, err.Error())
	}

	return &types.AdminTaskInfoReply{
		BasicReply: types.BasicReply(apiErr.Success),
		Task: types.Task{
			Id:          TaskInfoReply.Task.Id,
			UserId:      TaskInfoReply.Task.UserId,
			Title:       TaskInfoReply.Task.Title,
			Type:        TaskInfoReply.Task.Type,
			Text:        TaskInfoReply.Task.Text,
			QuestionUrl: TaskInfoReply.Task.QuestionUrl,
			AnswerUrl:   TaskInfoReply.Task.AnswerUrl,
			Bonus:       TaskInfoReply.Task.Bonus,
			State:       TaskInfoReply.Task.State,
		},
	}, nil
}
