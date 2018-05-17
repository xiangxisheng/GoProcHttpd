package http

import (
    "time"
    "net/http"
    "../../package/db"
)

type OutputData struct {
    Time float64
    Path string`json:"Path"`
    RemoteAddr string
    //ContentLength int64
    Message string
    Table db.Table
}

func GetOutputData(req *http.Request)OutputData {
    outputData := OutputData{}
    outputData.Time = float64(time.Now().UnixNano()) / (1000 * 1000 * 1000)
    outputData.Path = req.URL.Path
    outputData.RemoteAddr = req.RemoteAddr
    //outputData.ContentLength = req.ContentLength
    return outputData
}

