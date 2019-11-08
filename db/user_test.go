package db

import (
	"testing"
	"work-manager/pkg/utils"
)

const (
	userID   = uint(202201201709)
	realName = "马鑫"
	level    = uint16(2016)
	major    = "英语"
	uType    = uint8(0)

	account  = uint(202201201709)
	psssword = "maxinz"
	_gradeID = 2
)

func TestGetGrade(t *testing.T) {
	gradeID, err := GetGrade(level, major)
	if err != nil {
		t.Fatal(err)
	}

	if gradeID != uint(_gradeID) {
		t.Fatal("result is wrong")
	}

}

func TestCreateUser(t *testing.T) {
	passWord, _ := utils.MakePassword(psssword)
	_, err := CreateUser(userID, realName, passWord, level, major, uType)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetUser(t *testing.T) {
	if ok, err := GetUser(account, psssword); err != nil || !ok {
		t.Fatal("login faild")
	} else {
		t.Log("login successful")
	}
}
