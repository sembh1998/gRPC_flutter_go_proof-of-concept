package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnection struct {
	Conn *sql.DB
}

var mysqlSingleton *MysqlConnection

func getMySQLStringConnection() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	return dataBase
}

func GetMysqlConnection() *MysqlConnection {
	if mysqlSingleton == nil {
		conn, err := sql.Open("mysql", getMySQLStringConnection())
		if err != nil {
			panic(err)
		}
		mysqlSingleton = &MysqlConnection{Conn: conn}
		log.Println("Mysql connection created")
	}
	return mysqlSingleton
}
