package apiErr

// 本错误码参照了HTTP状态码的语义，方便识别错误类型

//分类	分类描述
//1**	信息，服务器收到请求，需要请求者继续执行操作
//2**	成功，操作被成功接收并处理
//3**	重定向，需要进一步的操作以完成请求
//4**	客户端错误，请求包含语法错误或无法完成请求
//5**	服务器错误，服务器在处理请求的过程中发生了错误

// Success 根据官方文档 0 代表成功
var Success = newError(0, "Success")

//// 200 OK
//var (
//	Success = newError(200, "Success")
//)

// 400 BAD REQUEST
var (
	InvalidParams      = newError(40001, "参数错误")
	PasswordIncorrect  = newError(40002, "密码错误")
	CheckActionUnknown = newError(40003, "未知的审核任务操作")
	FileUploadFailed   = newError(40004, "文件上传失败") //实际上也包括了文件打开上传后的关闭失败
	FileIsNotImage     = newError(40005, "文件不是图片")
	InvalidRequestBody = newError(40005, "json解析错误")
)

// 401 WITHOUT PERMISSION
var (
	NotLogin     = newError(40101, "用户未登录")
	InvalidToken = newError(40102, "无效的Token")
)

// 403 ILLEGAL OPERATION
var (
	PermissionDenied = newError(40301, "权限不足")
	IllegalOperation = newError(40302, "非法操作")
)

// 404 NOT FOUND
var (
	UserNotFound = newError(40401, "用户不存在")
	TaskNotFound = newError(40402, "任务不存在")
)

// 409 CONFLICT
var (
	UserAlreadyExist   = newError(40901, "用户已存在")
	TaskAlreadyChecked = newError(40904, "当前已检查完成")
)

// 500 INTERNAL ERROR
var (
	ServerInternal      = newError(50001, "服务器内部错误")
	GenerateTokenFailed = newError(50002, "生成Token失败")
)
