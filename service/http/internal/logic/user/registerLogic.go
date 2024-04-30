package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/common/utils"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterReply, err error) {
	logx.WithContext(l.ctx).Infof("注册: %v", req)

	//此处的基本业务逻辑都已经在handler中完成了，因为要处理文件。这里在为了维持统一性，分层清晰，借用一下登录的逻辑，但稍微做下区分
	// 调用rpc
	GetUserByNameReply, err := l.svcCtx.UserRpc.GetUserByName(l.ctx, &user.GetUserByNameRequest{
		Name: req.Username,
	})
	if rpcErr.Is(err, rpcErr.UserNotExist) {
		return nil, apiErr.UserNotFound
	} else if err != nil {
		logx.WithContext(l.ctx).Errorf("RegisterLogic.Register GetUserByName err: %v", err)
		return nil, apiErr.InternalError(l.ctx, err.Error())
	}

	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(GetUserByNameReply.Password), []byte(req.Password))
	if err != nil {
		return nil, apiErr.PasswordIncorrect
	}

	// 生成 token
	jwtToken, err := utils.CreateToken(
		GetUserByNameReply.Id,
		l.svcCtx.Config.UserAuth.AccessSecret,
		l.svcCtx.Config.UserAuth.AccessExpire,
	)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("RegisterLogic.Register CreateToken err: %v", err)
		return nil, apiErr.GenerateTokenFailed
	}

	return &types.RegisterReply{
		BasicReply: types.BasicReply(apiErr.Success),
		UserId:     GetUserByNameReply.Id,
		Token:      jwtToken,
	}, nil

}
