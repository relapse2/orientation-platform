package task

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
	"orientation-platform/service/rpc/task/types/task"
)

type TaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskInfoLogic {
	return &TaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskInfoLogic) TaskInfo(req *types.TaskInfoRequest) (resp *types.TaskInfoReply, err error) {
	logx.WithContext(l.ctx).Infof("获取任务信息: %v", req)

	//获取任务信息
	TaskInfoReply, err := l.svcCtx.TaskRpc.TaskInfo(l.ctx, &task.TaskInfoRequest{
		TaskId: req.TaskId,
	})
	if rpcErr.Is(err, rpcErr.TaskNotExist) {
		return nil, apiErr.TaskNotFound
	} else if err != nil {
		logx.WithContext(l.ctx).Errorf("获取任务信息失败: %v", err)
		return nil, apiErr.InternalError(l.ctx, err.Error())
	}

	return &types.TaskInfoReply{
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
