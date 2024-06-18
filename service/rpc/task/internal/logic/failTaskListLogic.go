package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/common/model"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type FailTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFailTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FailTaskListLogic {
	return &FailTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FailTaskListLogic) FailTaskList(in *task.Empty) (*task.FailTaskListReply, error) {
	var Tasks []model.Task
	//state为1是需要人工审核的
	if err := l.svcCtx.DBList.Mysql.Model(&model.Task{}).Where("state = ?", 1).Find(&Tasks).Error; err != nil {
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}
	if len(Tasks) == 0 {
		return nil, status.Error(rpcErr.TaskNotLoaded.Code, rpcErr.TaskNotLoaded.Message)
	}

	var TaskList []*task.TaskInfo
	for _, t := range Tasks {
		TaskList = append(TaskList, &task.TaskInfo{
			Id:          int64(t.ID),
			UserId:      t.UserId,
			Title:       t.TaskInfo.Title,
			Text:        t.TaskInfo.Text,
			Type:        t.TaskInfo.Type,
			Bonus:       t.TaskInfo.Bonus,
			State:       t.State,
			QuestionUrl: t.TaskInfo.ImageUrl,
			AnswerUrl:   t.AnswerUrl,
		})
	}

	return &task.FailTaskListReply{
		FailTaskList: TaskList,
	}, nil
}
