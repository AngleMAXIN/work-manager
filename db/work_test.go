package db

import (
	"testing"
	"time"
)

const (
	title     = "软件工程第一次测试1"
	_realName = "牛莉"
	creatorID = 33

	gradeID = 1

	Aname  = "刘生322"
	Atitle = "大作业21"
)

func TestCreateOneHomeWork(t *testing.T) {
	_endTime := time.Now()
	_startTime := time.Now()
	ok, err := CreateOneHomeWork(title, _realName, creatorID, _endTime, _startTime, gradeID)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("faild...")
	}
}

func TestCreateOneWork(t *testing.T) {
	ok, err := CreateOneWork(Aname, Atitle, creatorID, gradeID)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("faild...")
	}
}

func TestGetGradeWorkList(t *testing.T) {
	limit := 7
	offset := 0
	gradeID := 1
	_, err := GetGradeWorkList(limit, offset, gradeID)
	if err != nil {
		t.Fatal(err)
	}
}

func Benchmark_GetGradeWorkList_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetGradeWorkList(7, 0, 1)
		if err != nil {
			b.Fatal(err)
		}
	}

}

func TestGetWorkList(t *testing.T) {
	limit := 2
	offset := 0
	res, err := GetWorkList(limit, offset)
	if err != nil {
		t.Fatal(err)
	}

	for i, v := range res.Homeworks {
		t.Log(i, v)
	}
}

func TestCreateCommitToWork(t *testing.T) {
	comment := "不错呦"
	score := 34
	id := 1
	ok, err := CreateCommitToWork(comment, score, id)
	if err != nil || !ok {
		t.Fatal("faild...", err)
	}
}
