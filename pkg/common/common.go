package common

// RegisterBody 注册请求体
type RegisterBody struct {
	UserID   uint   `json:"user_id" binding:"required"`
	RealName string `json:"real_name" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Major    string `json:"major" binding:"required"`
	Level    uint16 `json:"level" binding:"required"`
	UType    uint8  `json:"u_type" binding:"required"`
}

// LoginBody 登录请求体
type LoginBody struct {
	AccountID uint   `json:"account_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
