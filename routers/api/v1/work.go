package v1

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"work-manager/db"
	"work-manager/pkg/common"
	"work-manager/pkg/e"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @Summary 教师端或是学生端获取作业列表
// @Produce  json
// @Param offset query int true "跳过的数量"
// @Param limit query  int  true "一页返回的数量"
// @Success 200 {object} ResBody test "{"code":200,"data":{"count":1,"homeworks":[{"homework_id":1,"belong_class":1,"creator_id":33,"title":"软件工程第一次测试1","creator":"牛莉","create_time":"2019-11-10T13:54:12+08:00","start_time":"2019-11-10T13:54:12+08:00","end_time":"2019-11-10T13:54:12+08:00"},null]},"massage":"ok"}"
// @Router /api/v1/homework [GET]
func GetWorkList(c *gin.Context) {
	offset, err1 := strconv.Atoi(c.Query("offset"))
	limit, err := strconv.Atoi(c.Query("limit"))
	code := 200
	res := &common.HomeWorkList{}
	if err != nil || err1 != nil {
		code = e.INVALID_PARAMS
		goto end
	}
	if res, err = db.GetWorkList(limit, offset); err != nil {
		code = e.ERROR
		goto end
	}
end:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"massage": e.GetMsg(code),
		"data":    res,
	})
}

// GetOneWorkResult 获取某一个作业的结果
// func OneWorkResult(c *gin.Context) {}

// GetGradeWorkList 获取一个班的所有作业
func GetGradeWorkList(c *gin.Context) {
	code := 200
	if _, err := c.Cookie("wm_login"); err != nil {
		c.JSON(code, gin.H{
			"code": 301,
			"msg":  "用户未登录",
			"data": nil,
		})
	}
	offset, err1 := strconv.Atoi(c.Query("offset"))
	limit, err := strconv.Atoi(c.Query("limit"))

	gradeID := c.Query("grade_id")
	homeworkID := c.Query("homework_id")
	res := &common.WorkList{}
	if err != nil || err1 != nil {
		code = e.INVALID_PARAMS
		goto end
	}
	if res, err = db.GetGradeWorkList(limit, offset, gradeID, homeworkID); err != nil {
		fmt.Println(err)
		code = e.ERROR
		goto end
	}
end:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"massage": e.GetMsg(code),
		"data":    res,
	})
}

// PutComment 批改作业
func PutComment(c *gin.Context) {
	var err error
	code := 200
	if _, err = c.Cookie("wm_login"); err != nil {
		c.JSON(code, gin.H{
			"code": 301,
			"msg":  "用户未登录",
			"data": nil,
		})
	}
	pwc := &common.PutCommentBody{}
	if err = c.ShouldBindJSON(pwc); err != nil {
		code = e.INVALID_PARAMS
		goto end
	}

	if _, err = db.CreateCommitToWork(pwc.Comment, pwc.Score, pwc.WorkID); err != nil {
		code = e.ERROR
		goto end
	}
end:
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}

// CreateHomeWork 发布作业
func CreateHomeWork(c *gin.Context) {

}

// PostHomeWork 提交作业
func PostHomeWork(c *gin.Context) {
	var err error
	code := 200
	if _, err = c.Cookie("wm_login"); err != nil {
		c.JSON(code, gin.H{
			"code": 301,
			"msg":  "用户未登录",
			"data": nil,
		})
	}
	pwb := &common.PostWorkBody{}
	f, err := c.FormFile("upload_work_file")
	if err != nil {
		fmt.Println(err)
		code = e.ERROR_FILE_TOO_BIG
		goto end
	} else {
		filename := filepath.Base(f.Filename)
		workDir := fmt.Sprintf("%s%d", common.WorkBaseDir, pwb.HomeworkID)
		// 判断文件夹是否存在
		// 不存在就创建，存在就直接把文件写进去
		err := os.Mkdir(workDir, os.ModePerm)

		if err = c.SaveUploadedFile(f, workDir+filename); err != nil {
			fmt.Println(err)
			code = e.ERROR
			goto end
		}

		pwb.Title = f.Filename
		s := sessions.Default(c)
		u, _ := s.Get("user_info").(common.UserBody)
		pwb.Creator = u.RealName
		pwb.CreatorID = u.UserID
		pwb.GradeID = u.GradeID
		pwb.HomeworkID, _ = strconv.Atoi(c.PostForm("homework_id"))
	}
	if _, err = db.CreateOneWork(pwb); err != nil {
		code = e.ERROR
		goto end
	}
end:
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}
