package dbUtil

import (
	"github.com/gohouse/gorose"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"HFish/utils/conf"
)

var engin *gorose.Engin

func init() {
	var err error

	dbType := conf.Get("admin", "db_type")
	dbStr := conf.Get("admin", "db_str")
	dbMaxOpen := conf.GetInt("admin", "db_max_open")

	if dbType == "sqlite" {
		engin, err = gorose.Open(&gorose.Config{Driver: "sqlite3", Dsn: dbStr, SetMaxOpenConns: dbMaxOpen})

		if err != nil {
			fmt.Println(err)
		}
	} else if dbType == "mysql" {
		engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: dbStr, SetMaxOpenConns: dbMaxOpen})

		if err != nil {
			fmt.Println(err)
		}
	}
}

func DB() gorose.IOrm {
	return engin.NewOrm()
}
