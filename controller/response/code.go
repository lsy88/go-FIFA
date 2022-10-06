package response

type StatusCode int64

const (
	Success              StatusCode = 1000
	ERR_Invalid_Params   StatusCode = 1001
	ERR_User_Exist       StatusCode = 1002
	ERR_User_NotExist    StatusCode = 1003
	ERR_Invalid_Password StatusCode = 1004
	ERR_Server_Busy      StatusCode = 1005
	
	ERR_Invalid_Token      StatusCode = 1006
	ERR_Invalid_AuthFormat StatusCode = 1007
	ERR_NotLogin           StatusCode = 1008
	ERR_Token_NotExist     StatusCode = 1009
	ERR_NotRegister        StatusCode = 1010
)

var MessageFlag = map[StatusCode]string{
	Success:              "success",
	ERR_Invalid_Params:   "请求参数错误",
	ERR_User_Exist:       "用户名已存在",
	ERR_User_NotExist:    "用户不存在",
	ERR_Invalid_Password: "密码错误",
	ERR_Server_Busy:      "服务繁忙",
	
	ERR_Invalid_Token:      "无效的Token",
	ERR_Token_NotExist:     "请求头缺少Token",
	ERR_Invalid_AuthFormat: "认证格式有误",
	ERR_NotLogin:           "未登录",
	ERR_NotRegister:        "注册失败",
}

func (s StatusCode) Decode() string {
	message, ok := MessageFlag[s]
	if ok {
		//存在直接返回
		return message
	}
	//不存在返回服务忙
	return MessageFlag[ERR_Server_Busy]
}
