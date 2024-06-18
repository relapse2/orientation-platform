package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/common/model"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskInfoLogic {
	return &TaskInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskInfoLogic) TaskInfo(in *task.TaskInfoRequest) (*task.TaskInfoReply, error) {
	// 准备模型
	t := &model.Task{}

	// 查询数据
	err := l.svcCtx.DBList.Mysql.Where("id = ?", in.TaskId).First(t).Error

	if err == gorm.ErrRecordNotFound {
		return nil, status.Error(rpcErr.TaskNotExist.Code, rpcErr.TaskNotExist.Message)
	} else if err != nil {
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	return &task.TaskInfoReply{
		Task: &task.TaskInfo{
			Id:          int64(t.ID),
			UserId:      t.UserId,
			Title:       t.TaskInfo.Title,
			Type:        t.TaskInfo.Type,
			Text:        t.TaskInfo.Text,
			QuestionUrl: t.TaskInfo.ImageUrl,
			AnswerUrl:   t.AnswerUrl,
			Bonus:       t.TaskInfo.Bonus,
			State:       t.State,
		},
	}, nil
}
