package rpcErr

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//这里使用了grpc的status来处理rpc错误，但code进行了自定义以知晓具体错误

const (
	DataBaseErrorCode codes.Code = iota + 1000
	CacheErrorCode
	PasswordEncryptFailedCode
	MQErrorCode

	ConvertStringCode
)

const (
	UserAlreadyExistCode codes.Code = iota + 2000
	StuNotLoadedCode
	StuAlreadyLoadedCode
	TaskNotLoadedCode
	UserNotExistCode
	TaskNotExistCode
)

const (
	CommentNotExistCode codes.Code = iota + 3000
)

var errCodeMap = map[codes.Code]string{
	DataBaseErrorCode:         "数据库错误",
	CacheErrorCode:            "缓存错误",
	PasswordEncryptFailedCode: "密码加密失败",
	MQErrorCode:               "消息队列错误",

	StuNotLoadedCode:     "用户对应学生信息未导入",
	StuAlreadyLoadedCode: "用户对应学生信息已导入",
	TaskNotLoadedCode:    "学校对应任务信息未导入",

	UserAlreadyExistCode: "用户已存在",
	UserNotExistCode:     "用户不存在",
	TaskNotExistCode:     "任务不存在",

	CommentNotExistCode: "评论不存在",
	ConvertStringCode:   "经纬度转换类型错误",
}

var (
	StuNotLoaded          = NewRpcError(StuNotLoadedCode)
	StuAlreadyLoaded      = NewRpcError(StuAlreadyLoadedCode)
	TaskNotLoaded         = NewRpcError(TaskNotLoadedCode)
	UserAlreadyExist      = NewRpcError(UserAlreadyExistCode)
	DataBaseError         = NewRpcError(DataBaseErrorCode)
	CacheError            = NewRpcError(CacheErrorCode)
	MQError               = NewRpcError(MQErrorCode)
	PassWordEncryptFailed = NewRpcError(PasswordEncryptFailedCode)
	UserNotExist          = NewRpcError(UserNotExistCode)
	TaskNotExist          = NewRpcError(TaskNotExistCode)
	CommentNotExist       = NewRpcError(CommentNotExistCode)
	ConvertString         = NewRpcError(ConvertStringCode)
)

type RpcError struct {
	Code    codes.Code `json:"code"`
	Message string     `json:"message"`
}

func NewRpcError(code codes.Code) *RpcError {
	return &RpcError{
		Code:    code,
		Message: errCodeMap[code],
	}
}

func Is(err error, target *RpcError) bool {

	if err == nil {
		return false
	}
	s, _ := status.FromError(err)

	return s.Code() == target.Code
}

func (e *RpcError) Error() string {
	return e.Message
}
