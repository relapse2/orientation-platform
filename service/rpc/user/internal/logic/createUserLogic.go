package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/common/model"
	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreatUserReply, error) {
	tx := l.svcCtx.DBList.Mysql.Begin()

	//检查是否有匹配的学生信息已经导入,以及是否已经存在此学生
	result := &model.Student{}
	if err := tx.Model(&model.Student{}).Where("id_card = ?", in.IdCard).Limit(1).Find(result).Error; err != nil {
		tx.Rollback()
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}
	if result.ID == 0 {
		return nil, status.Error(rpcErr.StuNotLoaded.Code, rpcErr.StuNotLoaded.Message)
	}

	if result.IfVerified == true {
		tx.Rollback()
		return nil, status.Error(rpcErr.UserAlreadyExist.Code, rpcErr.UserAlreadyExist.Message)
	} else {
		tx.Model(result).Update("IfVerified", true)
	}

	// 密码加密
	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return nil, status.Error(rpcErr.PassWordEncryptFailed.Code, err.Error())
	}

	//转化经纬度
	Longitude, err := strconv.ParseFloat(in.Longitude, 64)
	if err != nil {
		return nil, status.Error(rpcErr.ConvertString.Code, err.Error())
	}

	Latitude, err := strconv.ParseFloat(in.Latitude, 64)
	if err != nil {
		return nil, status.Error(rpcErr.ConvertString.Code, err.Error())
	}

	//判断性别
	// 获取第 17 位数字
	genderDigit := in.IdCard[16]

	// 根据第 17 位数字判断性别
	var sex int64
	if genderDigit%2 == '1' {
		sex = 1
	} else {
		sex = 0
	}

	// 准备数据
	newUser := &model.User{
		Username: in.Name,
		Password: string(password),
		Collage:  in.Collage,
		Sex:      sex,
		IdCard:   in.IdCard,
		RegAddress: model.Location{
			Address:   in.Address,
			Latitude:  Latitude,
			Longitude: Longitude,
		},
		AvatarUrl: in.AvatarUrl,
	}

	// 插入数据
	if err := tx.Create(newUser).Error; err != nil {
		tx.Rollback()
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	return &user.CreatUserReply{
		Id: int64(newUser.ID),
	}, nil
}
