package db

import (
	"database/sql"
	"errors"
	"time"
	"work-manager/pkg/common"
)

const (
	// createHomoWork  老师创建作业
	createHomoWork = `insert into wm_homework
					    (title, creator_id, creator, create_time, start_time, end_time, belong_class) 
					  values (?,?,?,?,?,?,?);`

	// getHomeWorkOfGrade 获取某一个班的已提交的作业
	getWorkOfGrade = `select 
						   id, creator, title, creator_id, comment, score, upload_time, grade_id from wm_work 
					  where homework_id = ? limit ? offset ?;`

	// getCountWStr 获取某一个班的已提交的作业数量
	getCountWStr = `select count(*) from wm_work where homework_id = ?;`

	// getHomeworkListStr 获取所有布置的作业
	getHomeworkListStr = `select * from wm_homework limit ? offset ?;`

	// getCountHwStr 获取所有布置的作业的数量
	getCountHwStr = `select count(*) from wm_homework;`

	// createOneWork 创建一个作业
	createOneWork = `insert into wm_work (creator, title, upload_time, creator_id, grade_id, homework_id) values (?,?,?,?,?,?) on duplicate key update title=values(title);`

	// createCommentStr 老师评价作业
	createCommentStr = `update wm_work set comment = ? ,score = ? where id = ?;`

	// deleteWorklistStr 删除有关的所有已提交的作业
	// deleteWorklistStr = `delete from wm_work where homework_id = ?;`

	// deleteHomeWorkStr 删除布置的作业
	// deleteHomeWorkStr = `delete from wm_homework where id = ?;`

	// deleteOneWorkStr 删除某一个提交过的作业记录
	deleteOneWorkStr = `delete from wm_work where id = ?;`
)

// DeleteWork 删除提交的作业
func DeleteWork(workID string) (bool, error) {
	var (
		err  error
		stmt *sql.Stmt
	)
	if stmt, err = dbConn.Prepare(deleteOneWorkStr); err != nil {
		return false, err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(workID); err != nil {
		return false, err
	}
	return true, nil
}

// DeleteListHomeWork 删除已提交的作业
// func DeleteListHomeWork(homeWorkID int) (bool, error) {
// 	var (
// 		err  error
// 		stmt *sql.Stmt
// 	)
// 	if stmt, err = dbConn.Prepare(deleteWorklistStr); err != nil {
// 		return false, err
// 	}
// 	defer stmt.Close()

// 	if _, err = stmt.Exec(homeWorkID); err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// CreateOneHomeWork 老师创建作业
func CreateOneHomeWork(title, realName string, creatorID, gradeID int, endTime, startTime time.Time) (bool, error) {
	var (
		err  error
		stmt *sql.Stmt
	)
	if startTime.After(endTime) {
		e := "start time must before end time."
		return false, errors.New(e)
	}
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
func CreateOneWork(pwb *common.PostWorkBody) (bool, error) {
	var (
		err  error
		stmt *sql.Stmt
	)
	// 存在就更新，
	// 不存在就创建
	if stmt, err = dbConn.Prepare(createOneWork); err != nil {
		return false, err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(pwb.Creator, pwb.Title, time.Now(), pwb.CreatorID, pwb.GradeID, pwb.HomeworkID); err != nil {
		return false, err
	}
	return true, nil

}

// GetGradeWorkList 获取某一个专业的作业列表
func GetGradeWorkList(limit, offset int, homeworkID string) (*common.WorkList, error) {
	rows, err := dbConn.Query(getWorkOfGrade, homeworkID, limit, offset)
	if err != nil {
		return nil, err
	}
	count := 0
	if err = dbConn.QueryRow(getCountWStr, homeworkID).Scan(&count); err != nil {
		return nil, err
	}
	if limit >= count {
		limit = count
	}

	workList := &common.WorkList{
		Works: make([]*common.OneWork, limit),
		Count: count,
	}
	count = 0
	for rows.Next() {
		w := common.OneWork{}
		// id, creator, title, creator_id, comment, score, upload_time, grade_id
		if err = rows.Scan(&w.ID, &w.Creator, &w.Title, &w.CreatorID, &w.Comment, &w.Score, &w.UploadTime, &w.GradeID); err != nil {
			return nil, err
		}
		// workList.Works = append(workList.Works, &w)
		workList.Works[count] = &w
		count++
	}
	return workList, nil
}

// GetWorkList 获取所有的布置的作业信息列表，两种情况，学生和老师
func GetWorkList(limit, offset int) (*common.HomeWorkList, error) {
	rows, err := dbConn.Query(getHomeworkListStr, limit, offset)
	if err != nil {
		return nil, err
	}
	count := 0
	if err = dbConn.QueryRow(getCountHwStr).Scan(&count); err != nil {
		return nil, err
	}
	workList := &common.HomeWorkList{
		Homeworks: make([]*common.HomeWork, limit),
		Count:     count,
	}
	count = 0
	for rows.Next() {
		hw := common.HomeWork{}
		if err = rows.Scan(&hw.ID, &hw.Title, &hw.CreatorID, &hw.Creator, &hw.CreateTime, &hw.StartTime, &hw.EndTime, &hw.BelongClass); err != nil {
			return nil, err
		}
		workList.Homeworks[count] = &hw
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
