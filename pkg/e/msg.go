package e

// MsgFlags 返回题消息提示内容
var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	NOT_LOGIN:            "用户未登录",
	INVALID_PARAMS:       "请求参数错误",
	ERROR:                "系统错误",
	ERROR_EXIST_USER:     "该用户已注册",
	ERROR_NOT_EXIST_USER: "用户不存在",
	ERROR_USER_PASSWORD:  "密码错误",
	ERROR_FILE:           "文件错误",
}

// GetMsg 获取消息信息体
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

// {
// 	#返回状态码
// 	code:integer,

// 	#返回信息描述
// 	message:string,

// 	#返回值
// 	data:object
// }
