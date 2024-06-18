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

type TaskSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskSetLogic {
	return &TaskSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskSetLogic) TaskSet(in *task.TaskSetRequest) (*task.Empty, error) {
	tx := l.svcCtx.DBList.Mysql.Begin()
	// 准备数据
	newTask := &model.SetUpTask{
		Collage:        in.Collage,
		Title:          in.Title,
		Text:           in.Text,
		ImageUrl:       in.QuestionUrl,
		ReferAnswerUrl: in.ReferAnswerUrl,
		Bonus:          in.Bonus,
		Type:           in.Type,
	}

	// 插入数据
	if err := tx.Create(newTask).Error; err != nil {
		tx.Rollback()
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	return &task.Empty{}, nil
}
