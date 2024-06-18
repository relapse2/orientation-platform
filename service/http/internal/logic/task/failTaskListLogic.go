package task

import (
	"context"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/service/rpc/task/taskclient"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FailTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFailTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FailTaskListLogic {
	return &FailTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FailTaskListLogic) FailTaskList(req *types.FailTaskListRequest) (resp *types.FailTaskListReply, err error) {
	logx.WithContext(l.ctx).Infof("FailTaskList req: %+v", req)

	// 获取任务列表
	FailTaskListReply, err := l.svcCtx.TaskRpc.FailTaskList(l.ctx, &taskclient.Empty{})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("FailTaskList err: %+v", err)
		return nil, apiErr.InternalError(l.ctx, err.Error())
	}

	// 封装返回体
	resp = &types.FailTaskListReply{}
	resp.BasicReply = types.BasicReply(apiErr.Success)

	for _, t := range FailTaskListReply.FailTaskList {

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
