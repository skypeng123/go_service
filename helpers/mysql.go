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

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", system.GetConfiguration().MysqlDSN)
	if err != nil {
		seelog.Critical(err.Error())
		return nil, err
	}

	DB = db
	db.SetMaxIdleConns(system.GetConfiguration().MysqlMaxIdleConns)
	db.SetMaxOpenConns(system.GetConfiguration().MysqlMaxOpenConns)

	err = db.Ping()
	if err != nil {
		seelog.Critical(err.Error())
		return nil, err
	}
	return db, err
}
