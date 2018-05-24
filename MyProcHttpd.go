/*
go get -u github.com/tidwall/gjson
go get -u github.com/go-sql-driver/mysql
*/
package main

import (
    "os"
    "fmt"
    "./package/http"
)

func main() {
    fmt.Print("【HTTP开放接口】服务器")
    arg_num := len(os.Args)
    if false {
        fmt.Printf("\nthe num of input is %d", arg_num)
    }
    http.ListenAndServe()
}

