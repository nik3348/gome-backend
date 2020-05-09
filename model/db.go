package model

import (
	"GoMe/conf"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"strings"
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

// TODO: Only works if database isn't there
func InitData() {
	loadSQL(conf.Init)
	loadSQL(conf.CoursesSQL)
	loadSQL(conf.UsersSQL)
}

func loadSQL(file string) {
	data, err := ioutil.ReadFile(file) // Reading SQL file
	if err != nil {
		log.Fatal(err)
	}

	queries := strings.Split(string(data), ";\r\n")
	for i := 0; i < len(queries)-1; i++ {
		_, err := DB.Query(queries[i])
		fmt.Print(i)
		fmt.Println(":" + queries[i] + "\n")
		if err != nil {
			panic(err.Error())
		}
	}
}
