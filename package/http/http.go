package http

import (
    "fmt"
    "net/http"
    //"github.com/tidwall/gjson"
    "../../package/db"
)

var err error
var http_handler http.Handler

func checkErr(err error) {
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
}

func ListenAndServe() {

    db.Open("mysql", "feieryun:feieryun@tcp(3324.mysql.firadio.net:3324)/firadio_yun?charset=utf8")
    defer db.Close()

    //http请求处理
    http.Handle("/css/", http.FileServer(http.Dir("template")))
    http.Handle("/js/", http.FileServer(http.Dir("template")))
    http.HandleFunc("/proc", HandlerProc)
    if false{
        http.HandleFunc("/", HandlerRoot)
    } else {
        http.Handle("/", http.FileServer(http.Dir("template/html")))
    }

    //绑定监听地址和端口
    listen_sockets := "0.0.0.0:3380"
    fmt.Printf("http.ListenAndServe At %s\n", listen_sockets)
    err = http.ListenAndServe(listen_sockets, http_handler)
    checkErr(err)

}

