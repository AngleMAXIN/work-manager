package v1

import "github.com/gin-gonic/gin"

// GetWorkList 教师端或是学生端获取作业列表
func GetWorkList(c *gin.Context) {}

// GetOneWorkResult 获取某一个作业的结果
func GetOneWorkResult(c *gin.Context) {}

// GetGradeWorkList 获取一个班的所有作业
func GetGradeWorkList(c *gin.Context) {}

// CreateCommit 批改作业		
func CreateCommit(c *gin.Context) {}

// CreateToDoWork 发布作业
func CreateToDoWork(c *gin.Context) {}

// PushHomeWork 提交作业
func PushHomeWork(c *gin.Context) {}
