package user

import (
	"context"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/service/rpc/user/types/user"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoadStuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoadStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoadStuLogic {
	return &AdminLoadStuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoadStuLogic) AdminLoadStu(req *types.AdminLoadStuRequest) (resp *types.AdminLoadStuReply, err error) {
	logx.WithContext(l.ctx).Infof("导入学生信息: %v", req)

	var studentsValues []*user.Student
	for _, studentPtr := range req.Student {
		studentsValues = append(studentsValues, &user.Student{
			Name:    studentPtr.Name,
			Collage: studentPtr.Collage,
			IdCard:  studentPtr.IdCard,
		})
	}

	// 调用rpc
	_, err = l.svcCtx.UserRpc.AdminLoadStu(l.ctx, &user.AdminLoadStuRequest{
		Students: studentsValues,
	})

	if rpcErr.Is(err, rpcErr.UserAlreadyExist) {
		return nil, apiErr.UserAlreadyExist
	} else if err != nil {
		logx.WithContext(l.ctx).Errorf("调用rpc AdminLoadStu 失败: %v", err)
		return nil, apiErr.InternalError(l.ctx, err.Error())
	}

	return &types.AdminLoadStuReply{
		BasicReply: types.BasicReply(apiErr.Success),
	}, nil
}
