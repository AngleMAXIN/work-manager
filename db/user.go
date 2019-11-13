package db

import (
	"database/sql"
	"time"
	"work-manager/pkg/common"
)

const (
	// createUserStr 创建用户
	createUserStr = `insert into wm_user
						(user_id, real_name, u_type, password, level, major, grade_id, create_time) 
					values (?,?,?,?,?,?,?,?);`
	//getUserStr 获取用户
	getUserStr = `select id, password, grade_id, real_name, u_type from wm_user where user_id = ? limit 1;`
	//getGradeStr 获取某一个班级
	getGradeStr       = `select grade_id from wm_grade where level = ? and major = ?;`
	checkUserExistStr = `select count(1) from wm_user where user_id = ? limit 1;`
)

// GetUser 获取单个用户
func GetUser(account uint) (*common.UserBody, error) {
	u := &common.UserBody{}
	if err := dbConn.QueryRow(getUserStr, account).Scan(&u.UserID, &u.PassWord, &u.GradeID, &u.RealName, &u.UType); err != nil {
		return nil, err
	}
	return u, nil
}

// CreateUser 创建用户
func CreateUser(createBody *common.RegisterBody) (bool, error) {
	var (
		err     error
		stmt    *sql.Stmt
		gradeID uint
	)

	if stmt, err = dbConn.Prepare(createUserStr); err != nil {
		return false, err
	}
	defer stmt.Close()

	if gradeID, err = GetGrade(createBody.Level, createBody.Major); err != nil {
		return false, err
	}

	if _, err = stmt.Exec(createBody.UserID, createBody.RealName, createBody.UType,
		createBody.PassWord, createBody.Level, createBody.Major, gradeID, time.Now()); err != nil {
		return false, err
	}
	return true, nil
}

// GetGrade 获取某一个班级
func GetGrade(level uint16, major string) (uint, error) {
	var gradeID uint

	if err := dbConn.QueryRow(getGradeStr, level, major).Scan(&gradeID); err != nil {
		return 0, err
	}
	return gradeID, nil
}

// CheckUserExist 判断用户是否存在
func CheckUserExist(userID uint) bool {
	var exites uint

	if err := dbConn.QueryRow(checkUserExistStr, userID).Scan(&exites); err != nil || exites < 1 {
		return false
	}
	return true
}
