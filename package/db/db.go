package db

import (
    "fmt"
    "os"
    "strings"
    "time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var err error
var sqlDB*sql.DB

func checkErr(err error, title string) bool {
    if err != nil {
        date := time.Now().Format("2006-01-02 15:04:05.123")
        var msg string = ErrorMsg(err.Error())
        fmt.Fprintf(os.Stderr, "\n[Error] %s %s %s", date, title, msg)
        return true
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    return false
}

func ErrorMsg(msg string) string {
    if strings.Contains(msg, "invalid connection") {
        msg = "无效的连接（可能空闲连接超时被服务器断开）"
    }
    return msg
}

func Open(driver string, iqn string) (error) {
    // Open database connection
    sqlDB, err = sql.Open(driver, iqn)
    if checkErr(err, "sql.Open") { return err }
    // 取得数据库【全局变量】
    GetGlobalVariables()
    // 开始防掉线定时器
    go Timer()
    return err
}

func Close() {
    sqlDB.Close()
}

