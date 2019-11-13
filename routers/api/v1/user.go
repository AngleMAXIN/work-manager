package v1

import (
	"fmt"
	"net/http"
	"work-manager/db"
	"work-manager/pkg/common"
	"work-manager/pkg/e"
	"work-manager/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ResBody 返回体
type ResBody struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// @Summary 用户注册
// @Produce  json
// @Param user_id body int true "用户学号"
// @Param real_name body string  true "姓名"
// @Param password body string  true "密码"
// @Param major body string true  "软件工程"
// @Param level body int  true "年级"
// @Param u_type body int  true "用户类型"
// @Success 200 {object} ResBody "{"code":200,"data":nil,"msg":"ok"}"
// @Failure 400 {object} ResBody "{"code":400,"data":nil,"msg":"请求参数错误"}"
// @Failure 10001 {object} ResBody "{"code":10001,"data":nil,"msg":"该用户已注册"}"
// @Failure 500 {object} ResBody "{"code":500,"data":nil,"msg":"系统错误"}"
// @Router /api/v1/register [POST]
func Register(c *gin.Context) {
	var err error
	code := 200
	regBody := &common.RegisterBody{}
	if err = c.ShouldBindJSON(regBody); err != nil {
		code = e.INVALID_PARAMS
		goto end
	}
	if isExist := db.CheckUserExist(regBody.UserID); isExist {
		code = e.ERROR_EXIST_USER
		goto end
	}
	regBody.PassWord, _ = utils.MakePassword(regBody.PassWord)
	if _, err = db.CreateUser(regBody); err != nil {
		code = e.ERROR
	}
end:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"massage": e.GetMsg(code),
		"data":    nil,
	})
}

// @Summary 用户登录
// @Produce  json
// @Param account_id body int true "用户学号"
// @Param password body string  true "密码"
// @Success 200 {object} ResBody "{"code":200,"data":nil,"msg":"ok"}"
// @Failure 10002 {object} ResBody "{"code":10002,"data":nil,"msg":"用户不存在"}"
// @Failure 10003 {object} ResBody "{"code":10003,"data":nil,"msg":"密码错误"}"
// @Router /api/v1/login [POST]
func Login(c *gin.Context) {
	var err error
	code := 200
	loginBody := &common.LoginBody{}
	u := &common.UserBody{}
	if err = c.ShouldBindJSON(loginBody); err != nil {
		code = e.INVALID_PARAMS
		goto end
	}
	if u, err = db.GetUser(loginBody.AccountID); err != nil {
		fmt.Println(err)
		code = e.ERROR_NOT_EXIST_USER
		goto end
	}
	if ok := utils.CheckPassword(loginBody.Password, u.PassWord); !ok {
		code = e.ERROR_USER_PASSWORD
		goto end
	} else {
		session := sessions.Default(c)
		session.Options(sessions.Options{
			MaxAge: 3600 * 24 * 7,
			Path:   "/",
		})
		session.Set("user_info", *u)
		session.Save()
	}
end:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"massage": e.GetMsg(code),
		"data":    nil,
	})

}

// GradeList 班级信息
func GradeList(c *gin.Context) {}
