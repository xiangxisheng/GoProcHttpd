package http

import (
    "fmt"
    "os"
    "strings"
    "time"
    "net/http"
    //"github.com/tidwall/gjson"
    "../../package/db"
)

var err error
var http_handler http.Handler

func checkErr(err error, title string) bool {
    if err != nil {
        date := time.Now().Format("15:04:05.999")
        var msg string = ErrorMsg(err.Error())
        fmt.Fprintf(os.Stderr, "[Error] %s %s %s\n", date, title, msg)
        return true
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    return false
}

func ErrorMsg(msg string) string {
    if strings.Contains(msg, "bind: Only one usage") {
        msg = "端口被占用"
    }
    return msg
}

func ListenAndServe() {

    db.Open("mysql", "feieryun:feieryun@tcp(3324.mysql.firadio.net:3324)/firadio_yun?charset=utf8")
    defer db.Close()

    //http请求处理
    http.Handle("/css/", http.FileServer(http.Dir("template")))
    http.Handle("/js/", http.FileServer(http.Dir("template")))
    http.HandleFunc("/proc", HandlerProcInfo)
    http.HandleFunc("/proc/", HandlerProcCall)
    if false {
        http.HandleFunc("/", HandlerRoot)
    } else {
        http.Handle("/", http.FileServer(http.Dir("template/html")))
    }

    //绑定监听地址和端口
    listen_sockets := "0.0.0.0:3380"
    fmt.Printf("http.ListenAndServe At %s\n", listen_sockets)
    err = http.ListenAndServe(listen_sockets, http_handler)
    if checkErr(err, "http.ListenAndServe") { return }

}

