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
// @Success 200 {object} ResBody  {"code":200,"data":{"count":1,"homeworks":[{"homework_id":1,"belong_class":1,"creator_id":33,"title":"软件工程第一次测试1","creator":"牛莉","create_time":"2019-11-10T13:54:12+08:00","start_time":"2019-11-10T13:54:12+08:00","end_time":"2019-11-10T13:54:12+08:00"},null]},"massage":"ok"}
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

// @Summary 删除提交的作业
// @Produce  json
// @Param work_id query int true "自己提交作业id"
// @Param file_name query  string  true "作业文件名称"
// @Param homework_id query  int  true "当前布置的老师布置的作业id"
// @Success 200 {object} ResBody test "{"code":200,"data":null,"massage":"ok"}"
// @Success 500 {object} ResBody test "{"code":500,"data":null,"massage":"系统错误"}"
// @Router /api/v1/delete-work [DELETE]
// func DeleteWork(c *gin.Context) {
// 	workID := c.Query("work_id")
// 	workFile := c.Query("file_name")
// 	homeWorkID := c.Query("homework_id")

// 	var err error
// 	code := 200
// 	if _, err = db.DeleteWork(workID); err != nil {
// 		code = e.ERROR
// 		goto end
// 	}
// 	if err = os.Remove(fmt.Sprintf("%s/%s/%s", common.WorkBaseDir, homeWorkID, workFile)); err != nil {
// 		code = e.ERROR
// 		goto end
// 	}

// end:
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":    code,
// 		"massage": e.GetMsg(code),
// 		"data":    nil,
// 	})
// }

// // Summary 获取一个班的所有作业
// @Produce  json
// @Param offset query int true "跳过的数量"
// @Param limit query  int  true "一页返回的数量"
// @Param homework_id query  int  true "布置的作业id	"
// @Success 200 {object}  ResBody {"code":200,"data":{"Count":2,"Works":[{"work_id":1,"creator_id":0,"score":0,"grade_id":0,"homework_id":0,"creator":"","title":"方案","comment":"","upload_time":"2019-11-21T11:21:29+08:00"}]},"massage":"ok"} ResBody
// @Success 500 {object}  ResBody {"code":500,"data":null,"massage":"系统错误"}
// @Router /api/v1/grade-homework [GET]
func GetGradeWorkList(c *gin.Context) {
	var err error
	code := 200
	if _, err = c.Cookie("wm_login"); err != nil {
		c.JSON(code, gin.H{
			"code": 301,
			"msg":  "用户未登录",
			"data": nil,
		})
	}

	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	homeworkID := c.Query("homework_id")
	res := &common.WorkList{}

	if res, err = db.GetGradeWorkList(limit, offset, homeworkID); err != nil {
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

// @Summary 批改作业
// @Produce  json
// @Param comment body string true "评语"
// @Param score body int  true "分数"
// @Param work_id body int  true "某一个学生的作业id"
// @Success 200 {object} ResBody "{"code":200,"data":nil,"msg":"ok"}"
// @Failure 400 {object} ResBody "{"code":400,"data":nil,"msg":"请求参数错误"}"
// @Failure 301 {object} ResBody "{"code":10001,"data":nil,"msg":"用户未登录"}"
// @Failure 500 {object} ResBody "{"code":500,"data":nil,"msg":"系统错误"}"
// @Router /api/v1/comment-work [POST]
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

// @Summary 教师发布作业
// @Produce  json
// @Param level body int  true "年级"
// @Param major body string  true "专业"
// @Param title body string  true "布置作业的标题"
// @Param start_time body time  true "开始时间"
// @Param end_time body time  true "结束时间"
// @Success 200 {object} ResBody {"code":200,"data":nil,"msg":"ok"}
// @Failure 400 {object} ResBody {"code":400,"data":nil,"msg":"请求参数错误"}
// @Failure 301 {object} ResBody {"code":10001,"data":nil,"msg":"用户未登录"}
// @Failure 500 {object} ResBody {"code":500,"data":nil,"msg":"系统错误"}
// @Router /api/v1/create-homework [POST]
func CreateHomeWork(c *gin.Context) {
	var err error
	var gradeID uint
	code := 200
	if _, err = c.Cookie("wm_login"); err != nil {
		c.JSON(code, gin.H{
			"code": 301,
			"msg":  "用户未登录",
			"data": nil,
		})
	}
	hw := &common.CreateHomeWorkBody{}
	s := sessions.Default(c)
	u, _ := s.Get("user_info").(common.UserBody)
	hw.Creator = u.RealName
	hw.CreatorID = u.UserID

	if err = c.ShouldBindJSON(hw); err != nil {
		code = e.INVALID_PARAMS
		goto end
	}
	if gradeID, err = db.GetGrade(hw.Level, hw.Major); err != nil {
		code = e.INVALID_PARAMS
		goto end

	}
	if _, err = db.CreateOneHomeWork(hw.Title, hw.Creator, hw.CreatorID, int(gradeID), hw.EndTime, hw.StartTime); err != nil {
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

// @Summary 提交作业
// @Produce  application/x-www-form-urlencoded
// @Param homework_id body int true "布置的作业id"
// @Param title body string  true "作业文件标题"
// @Success 200 {object} ResBody "{"code":200,"data":nil,"msg":"ok"}"
// @Failure 10006 {object} ResBody "{"code":400,"data":nil,"msg":"文件错误"}"
// @Failure 301 {object} ResBody "{"code":10001,"data":nil,"msg":"用户未登录"}"
// @Failure 500 {object} ResBody "{"code":500,"data":nil,"msg":"系统错误"}"
// @Router /api/v1/post-work [POST]
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
	homeworkID := c.PostForm("homework_id")
	if err != nil {
		code = e.ERROR_FILE
		goto end
	} else {
		filename := filepath.Base(f.Filename)
		workDir := fmt.Sprintf("%s/%s/", common.WorkBaseDir, homeworkID)
		// 判断文件夹是否存在
		if _, err = os.Stat(workDir); err != nil {
			// 不存在就创建，存在就直接把文件写进去
			if !os.IsExist(err) {
				if err = os.Mkdir(workDir, os.ModePerm); err != nil {
					code = e.ERROR
					goto end
				}
			} else {
				code = e.ERROR
				goto end
			}
		}

		if err = c.SaveUploadedFile(f, workDir+filename); err != nil {
			code = e.ERROR
			goto end
		}

		// 从session中取出当前用户的信息
		pwb.Title = f.Filename
		s := sessions.Default(c)
		u, _ := s.Get("user_info").(common.UserBody)
		pwb.Creator = u.RealName
		pwb.CreatorID = u.UserID
		pwb.GradeID = u.GradeID
		pwb.HomeworkID, _ = strconv.Atoi(homeworkID)
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

// @Summary 获取文件
// @Produce  json
// @Param homework_id body int true "布置的作业id"
// @Param file_name body string  true "作业文件标题"
// @Success 200 {object} ResBody ，会下载下来一个文件
// @Router /api/v1/workfile [GET]
func GetWorkFile(c *gin.Context) {
	homeWorkID := c.Query("homework_id")
	fileName := c.Query("file_name")
	filePath := fmt.Sprintf("%s/%s/%s", common.WorkBaseDir, homeWorkID, fileName)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filePath)
}
