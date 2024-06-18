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

type SuccessTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSuccessTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuccessTaskLogic {
	return &SuccessTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SuccessTaskLogic) SuccessTask(in *task.AdminCheckTaskRequest) (*task.Empty, error) {
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

	newTask.State = 2

	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Save(&newTask).Error
	if err != nil {
		tx.Rollback()
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	tx.Commit()
	return &task.Empty{}, nil
}
