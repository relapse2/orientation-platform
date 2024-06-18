package task

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/utils"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
	"orientation-platform/service/rpc/task/taskclient"
)

type TaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskListLogic) TaskList(req *types.TaskListRequest) (resp *types.TaskListReply, err error) {
	logx.WithContext(l.ctx).Infof("GetTaskList req: %+v", req)

	// 获取登录用户id
	UserId, err := utils.GetUserIDFormToken(req.Token, l.svcCtx.Config.UserAuth.AccessSecret)
	if err != nil {
		UserId = 0
	}

	// 获取任务列表
	GetTaskListReply, err := l.svcCtx.TaskRpc.GetTaskList(l.ctx, &taskclient.GetTaskListRequest{
		UserId: UserId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetTaskList err: %+v", err)
		return nil, apiErr.InternalError(l.ctx, err.Error())
	}

	// 封装返回体
	resp = &types.TaskListReply{}
	resp.BasicReply = types.BasicReply(apiErr.Success)

	for _, t := range GetTaskListReply.TaskList {

		// 封装返回体
		resp.TaskList = append(resp.TaskList, types.Task{
			Id:          t.Id,
			UserId:      t.UserId,
			Title:       t.Title,
			Text:        t.Text,
			QuestionUrl: t.QuestionUrl,
			AnswerUrl:   t.AnswerUrl,
			Bonus:       t.Bonus,
			State:       t.State,
		})

	}
	return resp, nil
}
