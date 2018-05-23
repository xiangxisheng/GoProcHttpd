package http
import (
    //"fmt"
    "strings"
    "html/template"
    "net/http"
    "../../package/db"
)
//请求处理函数
func HandlerProcInfo(w http.ResponseWriter, req *http.Request) {
    //获取请求资源
    outputData := GetOutputData(req)
    outputData.Table, err = db.GetParamData()
    if checkErr(err, "handler.HandlerProcInfo") {
        outputData.Message = err.Error()
    }
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
    sUrlPath := req.URL.Path
    sUrlPath = strings.Replace(sUrlPath, "//", "/", -1)
    aUrlPath := strings.Split(sUrlPath, "/")
    if (len(aUrlPath) < 4) {
        outputData.Message = "提供的URL.Path不正确"
        WriteJSON(w, outputData)
        return
    }
    //取得【数据库】名称
    sDbName := aUrlPath[2]
    //取得【存储过程】名称
    sProcName := "proc_" + aUrlPath[3]
    //取得【存储过程】要求传递的参数
    aProcParam, err := db.GetProcParam(sDbName, sProcName)
    if checkErr(err, "HandlerProcCall.db.GetProcParam") {
        outputData.Message = err.Error()
        WriteJSON(w, outputData)
        return
    }
    req.ParseForm()
    params := make([]interface{}, len(aProcParam))
    for i, oProcParam := range aProcParam {
        if oProcParam.PARAMETER_NAME == "ipaddr" {
            params[i] = "123.123.123.123" //RemoteAddr
            continue
        }
        sFormValue := req.FormValue(oProcParam.PARAMETER_NAME)
        if len(sFormValue) > oProcParam.MAXIMUM_LENGTH {
            outputData.Message = "您输入的" + oProcParam.PARAMETER_NAME + "超出最大长度"
            WriteJSON(w, outputData)
            return
        }
        params[i] = sFormValue
    }
    outputData.Table, err = db.GetProcData(sDbName, sProcName, params)
    if checkErr(err, "HandlerProcCall.GetProcData") {
        outputData.Message = err.Error()
    }
    WriteJSON(w, outputData)
}

