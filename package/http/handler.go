package http
import (
    "fmt"
    "html/template"
    "net/http"
    "../../package/db"
)
//请求处理函数
func HandlerProc(w http.ResponseWriter, req *http.Request) {
    //获取请求资源
    outputData := GetOutputData(req)
    req.ParseForm()
    proc := req.FormValue("proc")
    if proc == "" {
        outputData.Table, err = db.GetParamData()
        //outputData.Message="请提供proc"
        WriteJSON(w, outputData)
        return
    }
    //outputData.Data = make(map[interface{}]interface{})
    //outputData.Data["123"] = 123;
    outputData.Table, err = db.GetProcData(proc)
    //fmt.Println(outputData)
    WriteJSON(w, outputData)
}

func HandlerRoot(w http.ResponseWriter, req *http.Request) {
    t, err := template.ParseFiles("template/html/index.html")
    if checkErr(err, "template.ParseFiles") { return }
    t.Execute(w, nil)
}

func HandlerProcCall(w http.ResponseWriter, req *http.Request) {
    //执行存储过程
    outputData := GetOutputData(req)
    fmt.Printf("HandlerProcCall=%s\r\n", req.URL)
    fmt.Printf("HandlerProcCall=%s\r\n", req.URL.RawQuery)

    req.ParseForm()
    proc := req.FormValue("proc")
    if proc == "" {
        outputData.Table, err = db.GetParamData()
        //outputData.Message="请提供proc"
        WriteJSON(w, outputData)
        return
    }
    //outputData.Data = make(map[interface{}]interface{})
    //outputData.Data["123"] = 123;
    outputData.Table, err = db.GetProcData(proc)
    //fmt.Println(outputData)
    WriteJSON(w, outputData)
}

