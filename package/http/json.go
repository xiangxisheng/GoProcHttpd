package http

import (
    "io"
    "encoding/json"
    "net/http"
)

func WriteJSON(w http.ResponseWriter, outputData OutputData) {
    //支持全域名访问，不安全，部署后需要固定限制为客户端网址
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if false{
        //outputData.Body=buf.String()
        w.Header().Set("content-type", "application/json")
    }
    outputJSON,err:=json.Marshal(outputData)
    if err != nil {
        //fmt.Println(string(outputJSON))
        return
    }
    io.WriteString(w, string(outputJSON))
}

