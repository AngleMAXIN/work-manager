package common

import "time"

const (
	// WorkFileDir 作业存储路径
	WorkFileDir = "./workFiles/"
)

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

// UserBody 用户体
type UserBody struct {
	UserID   int    `json:"user_id,omitempty"`
	GradeID  int    `json:"grade_id,omitempty"`
	UType    int    `json:"u_type,omitempty"`
	RealName string `json:"real_name,omitempty"`
	PassWord string `json:"pass_word,omitempty"`
}

// PostWorkBody 学生创建作业
type PostWorkBody struct {
	CreatorID  int    `json:"creator_id" binding:"required"`
	GradeID    int    `json:"grade_id" binding:"required"`
	HomeworkID int    `json:"homework_id" binding:"required"`
	Creator    string `json:"creator" binding:"required"`
	Title      string `json:"title" binding:"required"`
}

// PostHomeWorkBody 老师创建作业
type PostHomeWorkBody struct {
	BelongClass int    `json:"belong_class,omitempty"`
	CreatorID   int    `json:"creator_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Creator     string `json:"creator,omitempty"`
}

// PutCommentBody 批改作业
type PutCommentBody struct {
	Comment string `json:"comment,omitempty"`
	Score   int    `json:"score,omitempty"`
	WorkID  int    `json:"work_id,omitempty"`
}

// HomeWork 老师布置的作业
type HomeWork struct {
	ID          int       `json:"homework_id,omitempty"`
	BelongClass int       `json:"belong_class,omitempty"`
	CreatorID   int       `json:"creator_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Creator     string    `json:"creator,omitempty"`
	CreateTime  time.Time `json:"create_time,omitempty"`
	StartTime   time.Time `json:"start_time,omitempty"`
	EndTime     time.Time `json:"end_time,omitempty"`
}

// HomeWorkList 布置作业集合
type HomeWorkList struct {
	Count     int         `json:"count,omitempty"`
	Homeworks []*HomeWork `json:"homeworks,omitempty"`
}

// OneWork 单个作业
type OneWork struct {
	ID         int       `json:"work_id"`
	CreatorID  int       `json:"creator_id"`
	Score      int       `json:"score"`
	GradeID    int       `json:"grade_id"`
	HomeworkID int       `json:"homework_id"`
	Creator    string    `json:"creator"`
	Title      string    `json:"title"`
	Comment    string    `json:"comment"`
	UploadTime time.Time `json:"upload_time"`
}

// WorkList 提交作业集合
type WorkList struct {
	Count int
	Works []*OneWork
}
