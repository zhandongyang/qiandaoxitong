package app

var (
	// 服务器状态码
	OK                  = &Errno{Code: 200, Message: "OK"}
	InternalServerError = &Errno{Code: 500, Message: "服务器异常."}
	ErrBind             = &Errno{Code: 10002, Message: "将请求正文绑定到结构时发生错误."}

	// 用户错误
	ErrEncrypt    = &Errno{Code: 20101, Message: "加密用户密码时出错"}
	ErrPhoneExist = &Errno{Code: 20102, Message: "手机号已被注册"}
)
