package db

import (
	"testing"
	"time"
)

func TestCreateOneHomeWork(t *testing.T) {
	var (
		title      = "计算机科学第一周作业"
		_realName  = "商凯"
		creatorID  = 2
		gradeID    = 3
		_endTime   = time.Now().Add(time.Hour * 2)
		_startTime = time.Now().Add(time.Hour * 1)
	)

	_, err := CreateOneHomeWork(title, _realName, creatorID, gradeID, _endTime, _startTime)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetWorkList(t *testing.T) {
	limit := 2000
	offset := 3333
	_, err := GetWorkList(limit, offset)
	if err != nil {
		t.Fatal(err)
	}
}

// func TestCreateOneWork(t *testing.T) {
// 	ok, err := CreateOneWork(Aname, Atitle, creatorID, gradeID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !ok {
// 		t.Fatal("faild...")
// 	}
// }

// func TestGetGradeWorkList(t *testing.T) {
// 	limit := 7
// 	offset := 0
// 	gradeID := 1
// 	_, err := GetGradeWorkList(limit, offset, gradeID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func Benchmark_GetGradeWorkList_1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		_, err := GetGradeWorkList(7, 0, 1)
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 	}

// }

// func TestCreateCommitToWork(t *testing.T) {
// 	comment := "不错呦"
// 	score := 34
// 	id := 1
// 	ok, err := CreateCommitToWork(comment, score, id)
// 	if err != nil || !ok {
// 		t.Fatal("faild...", err)
// 	}
// }
