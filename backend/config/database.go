package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _"github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase(){
    connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"),
    )
    
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Database connection error: ", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Database unreachable: ", err)
    }

    log.Println("Database connected successfully.")

}

