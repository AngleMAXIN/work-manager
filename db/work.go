package db

import (
	"database/sql"
	"time"
)

const (
	// createHomoWork  老师创建作业
	createHomoWork = `insert into wm_homework
					    (title, creator_id, creator, create_time, start_time, end_time, belong_class) 
					  values (?,?,?,?,?,?,?);`

	// getHomeWorkOfGrade 获取某一个班的已提交的作业
	getWorkOfGrade = `select 
						   id, creator, title, creator_id, comment, score, upload_time, grade_id from wm_work 
					  where grade_id = ? limit ? offset ?;`

	// getCountWStr 获取某一个班的已提交的作业数量
	getCountWStr = `select count(*) from wm_work where grade_id = ?;`

	// getHomeworkListStr 获取所有布置的作业
	getHomeworkListStr = `select * from wm_homework limit ? offset ?;`

	// getCountHwStr 获取所有布置的作业的数量
	getCountHwStr = `select count(*) from wm_homework;`

	// createOneWork 创建一个作业
	createOneWork = `insert into wm_work (creator, title, upload_time, creator_id, grade_id) values (?,?,?,?,?);`

	// createCommentStr 老师评价作业
	createCommentStr = `update wm_work set comment = ? ,score = ? where id = ?;`
)

// HomeWork 老师布置的作业
type HomeWork struct {
	ID          int       `json:"homework_id"`
	BelongClass int       `json:"belong_class"`
	CreatorID   int       `json:"creator_id"`
	Title       string    `json:"title"`
	Creator     string    `json:"creator"`
	CreateTime  time.Time `json:"create_time"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

// HomeWorkList 布置作业集合
type HomeWorkList struct {
	Count     int
	Homeworks []*HomeWork
}

// OneWork 单个作业
type OneWork struct {
	ID         int       `json:"work_id"`
	CreatorID  int       `json:"creator_id"`
	Score      int       `json:"score"`
	GradeID    int       `json:"grade_id"`
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

// CreateOneHomeWork 老师创建作业
func CreateOneHomeWork(title string, realName string, creatorID int, endTime time.Time, startTime time.Time, gradeID int) (bool, error) {
	var (
		err  error
		stmt *sql.Stmt
	)
	if stmt, err = dbConn.Prepare(createHomoWork); err != nil {
		return false, err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(title, creatorID, realName, time.Now(), startTime, endTime, gradeID); err != nil {
		return false, err
	}
	return true, nil
}

// CreateOneWork 创建一个提交作业记录
func CreateOneWork(creator string, title string, creatorID int, gradeID int) (bool, error) {
	var (
		err  error
		stmt *sql.Stmt
	)

	if stmt, err = dbConn.Prepare(createOneWork); err != nil {
		return false, err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(creator, title, time.Now(), creatorID, gradeID); err != nil {
		return false, err
	}
	return true, nil

}

// GetGradeWorkList 获取某一个专业的作业列表
func GetGradeWorkList(limit, offset, gradeID int) (*WorkList, error) {
	rows, err := dbConn.Query(getWorkOfGrade, gradeID, limit, offset)
	if err != nil {
		return nil, err
	}
	count := 0
	if err = dbConn.QueryRow(getCountWStr, gradeID).Scan(&count); err != nil {
		return nil, err
	}
	workList := &WorkList{
		Works: make([]*OneWork, limit),
		Count: count,
	}
	count = 0
	for rows.Next() {
		one := OneWork{}
		if err = rows.Scan(&one.ID, &one.Creator, &one.Title, &one.CreatorID, &one.Comment, &one.Score, &one.UploadTime, &one.GradeID); err != nil {
			return nil, err
		}
		workList.Works[count] = &one
		count++
	}
	return workList, nil
}

// GetWorkList 获取所有的布置的作业信息列表，两种情况，学生和老师
func GetWorkList(limit, offset int) (*HomeWorkList, error) {
	rows, err := dbConn.Query(getHomeworkListStr, limit, offset)
	if err != nil {
		return nil, err
	}
	count := 0
	if err = dbConn.QueryRow(getCountHwStr).Scan(&count); err != nil {
		return nil, err
	}
	workList := &HomeWorkList{
		Homeworks: make([]*HomeWork, limit),
		Count:     count,
	}
	count = 0
	for rows.Next() {
		one := HomeWork{}
		if err = rows.Scan(&one.ID, &one.Title, &one.CreatorID, &one.Creator, &one.CreateTime, &one.StartTime, &one.EndTime, &one.BelongClass); err != nil {
			return nil, err
		}
		workList.Homeworks[count] = &one
		count++
	}
	return workList, nil
}

// CreateCommitToWork 批改作业
func CreateCommitToWork(comment string, score int, id int) (bool, error) {
	stmt, err := dbConn.Prepare(createCommentStr)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(comment, score, id); err != nil {
		return false, err
	}
	return true, nil
}

// GetWorkResult 获取某一个作业的结果
func GetWorkResult(id int) {

}
