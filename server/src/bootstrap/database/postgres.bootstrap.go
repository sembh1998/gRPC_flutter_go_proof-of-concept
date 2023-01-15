package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	Conn *sql.DB
}

var postgresSingleton *PostgresConnection

func getPostgresStringConnection() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
	)

	return dataBase
}

func GetPostgresConnection() *PostgresConnection {
	if postgresSingleton == nil {
		conn, err := sql.Open("postgres", getPostgresStringConnection())
		if err != nil {
			panic(err)
		}
		postgresSingleton = &PostgresConnection{Conn: conn}
		log.Println("Postgres connection created")
	}
	return postgresSingleton
}
