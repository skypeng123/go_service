package helpers

import (
	"database/sql"
	"github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"go_service/system"
)

var (
	DB *sql.DB
)

func InitDb() {
	var err error
	DB, err = sql.Open("mysql", system.GetConfiguration().MysqlDSN)
	if err != nil {
		seelog.Critical(err.Error())
		return
	}
	defer DB.Close()

	DB.SetMaxIdleConns(system.GetConfiguration().MysqlMaxIdleConns)
	DB.SetMaxOpenConns(system.GetConfiguration().MysqlMaxOpenConns)

	err = DB.Ping()
	if err != nil {
		seelog.Critical(err.Error())
		return
	}
}
