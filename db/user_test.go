package db

import (
	"testing"
	"work-manager/pkg/common"
	"work-manager/pkg/utils"
)

func TestGetGrade(t *testing.T) {
	var (
		level    = uint16(2017)
		major    = "计算机"
		_gradeID = 3
	)
	gradeID, err := GetGrade(level, major)
	if err != nil {
		t.Error(err)
	}

	if gradeID != uint(_gradeID) {
		t.Fatal("result is wrong")
	}
}

func TestCreateUser(t *testing.T) {
	registerBody := &common.RegisterBody{
		UserID:   uint(201601010309),
		PassWord: "maxinz",
		RealName: "马鑫",
		Major:    "计算机",
		Level:    2017,
		UType:    1,
	}
	registerBody.PassWord, _ = utils.MakePassword(registerBody.PassWord)
	_, err := CreateUser(registerBody)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUser(t *testing.T) {
	userID := uint(201601010309)
	if _, err := GetUser(userID); err != nil {
		t.Error(err)
	}
}

func TestCheckUserExist(t *testing.T) {
	userID := uint(201601010309)
	ok := CheckUserExist(userID)
	if !ok {
		t.Error("user is not exist")
	}
}
