package http
import (
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
    if proc==""{
        outputData.Table = db.GetParamData()
        //outputData.Message="请提供proc"
        WriteJSON(w, outputData)
        return
    }
    //outputData.Data = make(map[interface{}]interface{})
    //outputData.Data["123"] = 123;
    outputData.Table = db.GetProcData(proc)
    //fmt.Println(outputData)
    WriteJSON(w, outputData)
}

func HandlerRoot(w http.ResponseWriter, req *http.Request) {
    t, err := template.ParseFiles("template/html/index.html")
    checkErr(err)
    t.Execute(w, nil)
}

