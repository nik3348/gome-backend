package model

import (
	"GoMe/conf"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBConn() (*sql.DB, error) {
	dataSourceName := conf.USER + ":" + conf.PASSWORD + "@" + conf.PROTOCOL + "(" + conf.HOST + ":" + conf.PORT + ")/" + conf.DATABASE
	dbConn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = dbConn.Ping(); err != nil {
		return nil, err
	}

	DB = dbConn

	return dbConn, nil
}
