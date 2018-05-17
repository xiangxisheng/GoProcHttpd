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
    arg_num := len(os.Args)
    if false{
        fmt.Printf("the num of input is %d\n", arg_num)
    }
    http.ListenAndServe()
}

