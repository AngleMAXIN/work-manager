package db

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user = "root"
	loc  = "Asia/Shanghai"

	defaultHost     = "127.0.0.1"
	defaultPort     = "3306"
	defaultDB       = "work_manager"
	defaultPassword = "maxinz"
)

var (
	dbConn *sql.DB
)

func init() {

	host := defaultHost
	port := defaultPort
	db := defaultDB
	password := defaultPassword

	var err error
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s&parseTime=true",
		user, password, host, port, db, url.QueryEscape(loc))

	if dbConn, err = sql.Open("mysql", uri); err != nil {
		panic(err)
	}
}
