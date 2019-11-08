package v1

import (
	"net/http"
	"work-manager/db"
	"work-manager/pkg/common"
	"work-manager/pkg/e"
	"work-manager/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ResBody struct {
	Code    int         `json:"code,omitempty"`
	Massage string      `json:"massage,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// @Summary 用户注册
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {object} ResBody
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

// Login 用户登录
func Login(c *gin.Context) {
	var err error
	code := 200
	password := ""
	loginBody := &common.LoginBody{}
	if err = c.ShouldBindJSON(loginBody); err != nil {
		code = e.INVALID_PARAMS
		goto end
	}
	if password, err = db.GetUser(loginBody.AccountID); err != nil {
		code = e.ERROR_NOT_EXIST_USER
		goto end
	}
	// fmt.Println("----", password, loginBody.Password)
	if ok := utils.CheckPassword(loginBody.Password, password); !ok {
		code = e.ERROR_USER_PASSWORD
		goto end
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
