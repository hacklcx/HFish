package dbUtil

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gohouse/gorose"
	"HFish/utils/conf"
	"HFish/utils/log"
)

var engin *gorose.Engin

func init() {
	var err error

	dbType := conf.Get("admin", "db_type")
	dbStr := conf.Get("admin", "db_str")
	dbMaxOpen := conf.GetInt("admin", "db_max_open")
	dbMaxIdle := conf.GetInt("admin", "db_max_idle")

	if dbType == "sqlite" {
		engin, err = gorose.Open(&gorose.Config{Driver: "sqlite3", Dsn: dbStr, SetMaxOpenConns: dbMaxOpen, SetMaxIdleConns: dbMaxIdle})

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "Failed to connect to Sqlite database", err)
		}
	} else if dbType == "mysql" {
		engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: dbStr, SetMaxOpenConns: dbMaxOpen, SetMaxIdleConns: dbMaxIdle})

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "Failed to connect to Mysql database", err)
		}
	}
}

func DB() gorose.IOrm {
	return engin.NewOrm()
}
