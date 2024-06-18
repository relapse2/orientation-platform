package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/common/model"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type FailTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFailTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FailTaskLogic {
	return &FailTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FailTaskLogic) FailTask(in *task.AdminCheckTaskRequest) (*task.Empty, error) {
	// 开启事务
	tx := l.svcCtx.DBList.Mysql.Begin()
	var newTask model.Task
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", in.TaskId).First(&newTask).Error
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return nil, status.Error(rpcErr.TaskNotExist.Code, rpcErr.TaskNotExist.Message)
	} else if err != nil {
		tx.Rollback()
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	newTask.State = 0

	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Save(&newTask).Error
	if err != nil {
		tx.Rollback()
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	tx.Commit()
	return &task.Empty{}, nil
}
