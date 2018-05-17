package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var err error
var sqlDB*sql.DB

func checkErr(err error) {
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
}

func Open(driver string, iqn string) {
    // Open database connection
    sqlDB, err = sql.Open(driver, iqn)
    checkErr(err)
}

func Close() {
    sqlDB.Close()
}

