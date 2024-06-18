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

type TaskInitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskInitLogic {
	return &TaskInitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskInitLogic) TaskInit(in *task.TaskInitRequest) (*task.Empty, error) {

	tx := l.svcCtx.DBList.Mysql.Begin()
	//因为任务可以增加，set接口增加完任务之后时必须用一遍这个初始化增加任务，用户注册之后自动调用此rpc方法，TaskList里不加了
	//查询collage对应任务给user
	var SetUpTasks []model.SetUpTask

	if err := tx.Where("collage = ?", in.Collage).Find(&SetUpTasks).Error; err != nil {
		tx.Rollback()
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}
	if len(SetUpTasks) == 0 {
		return nil, status.Error(rpcErr.TaskNotLoaded.Code, rpcErr.TaskNotLoaded.Message)
	}

	for _, SetUpTask := range SetUpTasks {
		newTask := &model.Task{
			UserId:     in.UserId,
			State:      0,
			TaskInfoID: SetUpTask.ID,
			TaskInfo:   SetUpTask,
		}

		if err := tx.Create(newTask).Error; err != nil {
			tx.Rollback()
			return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
		}

		if err := tx.Commit().Error; err != nil {
			return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
		}
	}

	return &task.Empty{}, nil
}
