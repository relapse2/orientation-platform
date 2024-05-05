package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/common/model"

	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoadStuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoadStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoadStuLogic {
	return &AdminLoadStuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLoadStuLogic) AdminLoadStu(in *user.AdminLoadStuRequest) (*user.Empty, error) {
	tx := l.svcCtx.DBList.Mysql.Begin()

	var StuList []*model.Student

	for _, NewStudent := range in.Students {
		NewStu := &model.Student{
			Name:       NewStudent.Name,
			Collage:    NewStudent.Collage,
			IdCard:     NewStudent.IdCard,
			IfVerified: false,
		}

		//检查是否有学生信息已经导入
		result := &model.Student{}
		if err := tx.Model(&model.Student{}).Where("id_card = ?", NewStu.IdCard).Limit(1).Find(result).Error; err != nil {
			tx.Rollback()
			return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
		}
		if result.ID != 0 {
			return nil, status.Error(rpcErr.StuAlreadyLoaded.Code, rpcErr.StuAlreadyLoaded.Message)
		}

		StuList = append(StuList, NewStu)
	}

	// 插入数据
	if err := tx.Create(StuList).Error; err != nil {
		tx.Rollback()
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	return &user.Empty{}, nil
}
