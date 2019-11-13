package e

// MsgFlags 返回题消息提示内容
var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	INVALID_PARAMS:       "请求参数错误",
	ERROR:                "系统错误",
	ERROR_EXIST_USER:     "该用户已注册",
	ERROR_NOT_EXIST_USER: "用户不存在",
	ERROR_USER_PASSWORD:  "密码错误",
	ERROR_FILE_TOO_BIG:   "文件大小不得超过8M",
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
