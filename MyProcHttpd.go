/*
go get -u github.com/tidwall/gjson
go get -u github.com/go-sql-driver/mysql
*/
package main
import (
    "os"
    "bytes"
    "strconv"
    "fmt"
    "io"
    "time"
    "net/http"
    "encoding/json"
    "database/sql"
    "html/template"
    //"github.com/tidwall/gjson"
    _ "github.com/go-sql-driver/mysql"
)
var err error
var db *sql.DB
var http_handler http.Handler
func main() {
    arg_num := len(os.Args)
    if false{
        fmt.Printf("the num of input is %d\n", arg_num)
    }
    // Open database connection
    db, err = sql.Open("mysql", "feieryun:feieryun@tcp(3324.mysql.firadio.net:3324)/firadio_yun?charset=utf8")
    checkErr(err)
    defer db.Close()
    //http请求处理
    http.Handle("/css/", http.FileServer(http.Dir("template")))
    http.Handle("/js/", http.FileServer(http.Dir("template")))
    http.HandleFunc("/proc", http_handler_proc)
    if false{
        http.HandleFunc("/", http_handler_root)
    } else {
        http.Handle("/", http.FileServer(http.Dir("template/html")))
    }
    //绑定监听地址和端口
    listen_sockets := "0.0.0.0:3380"
    fmt.Printf("http.ListenAndServe At %s\n", listen_sockets)
    err = http.ListenAndServe(listen_sockets, http_handler)
    checkErr(err)
}
func GetParamData()Table{
    //sqlstr := "SELECT t1.SPECIFIC_SCHEMA,t1.SPECIFIC_NAME,t1.ORDINAL_POSITION,t1.PARAMETER_NAME,t1.DATA_TYPE,t1.CHARACTER_MAXIMUM_LENGTH,t1.NUMERIC_PRECISION FROM `information_schema`.`PARAMETERS` t1"
    //sqlstr := "SELECT t1.SPECIFIC_SCHEMA,t1.SPECIFIC_NAME FROM `information_schema`.`PARAMETERS` t1"
    var buf bytes.Buffer
    buf.WriteString("SELECT T.SPECIFIC_SCHEMA,T.SPECIFIC_NAME,MAX(PARAMETERS)PARAMETERS FROM (")
    buf.WriteString("SELECT t1.SPECIFIC_SCHEMA,t1.SPECIFIC_NAME,GROUP_CONCAT(t1.PARAMETER_NAME,'|',t1.DATA_TYPE,'|',IFNULL(t1.CHARACTER_MAXIMUM_LENGTH,''),IFNULL(t1.NUMERIC_PRECISION,'')ORDER BY ORDINAL_POSITION ASC)PARAMETERS FROM `information_schema`.`PARAMETERS` t1 GROUP BY t1.SPECIFIC_SCHEMA,t1.SPECIFIC_NAME")
    buf.WriteString(" UNION ")
    buf.WriteString("SELECT t2.ROUTINE_SCHEMA,t2.ROUTINE_NAME,NULL FROM `information_schema`.`ROUTINES` t2 WHERE t2.ROUTINE_TYPE='PROCEDURE'")
    buf.WriteString(")T GROUP BY T.SPECIFIC_SCHEMA,T.SPECIFIC_NAME ORDER BY T.SPECIFIC_NAME ASC")
    //fmt.Printf(buf.String())
    return GetTableBySql(buf.String())
}
func GetProcData(procName string) Table {
    sqlstr := "CALL `" + procName + "`()"
    return GetTableBySql(sqlstr)
}
func GetTableBySql(sqlstr string) Table {
    // Execute the query
    rows, err := db.Query(sqlstr)
    checkErr(err)
    var table Table=Table{}
    var table_rows [][]interface{}
    // Get column names
    columns, err := rows.Columns()
    checkErr(err)
    table.Columns = columns
    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))
    // rows.Scan wants '[]interface{}' as an argument, so we must copy the
    // references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }
    // Fetch rows
    for rows.Next() {
        // get RawBytes from data
        if true{
            err = rows.Scan(scanArgs...)
            checkErr(err)
        }
        // Now do something with the data.
        // Here we just print each column as a string.
        var sValues []interface{}
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            var value interface{}
            switch checkType(col){
                case "nil":
                    value = nil
                    break
                case "num":
                    value,err = strconv.Atoi(string(col))
                    //checkErr(err)
                    break
                case "string":
                    value = string(col)
                    break
                default:
                    value = col
                    break
            }
            sValues = append(sValues, value)
            if false{
                fmt.Println(columns[i], ": ", value)
            }
        }
        table_rows = append(table_rows, sValues)
    }
    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    rows.Close()
    table.Rows = table_rows
    return table
}
func checkType(rawBytes sql.RawBytes)string {
    if rawBytes==nil{
        return "nil"
    }
    var min byte=255
    var max byte=0
    for i, char := range rawBytes {
        if false{
            fmt.Println(i, ", ", char)
        }
        if char<min{
            min=char
        }
        if char>max{
            max=char
        }
    }
    if min<32{
        return "base64"
    }
    if min<48{
        return "string"
    }
    if max>57{
        return "string"
    }
    return "num"
}
func checkErr(err error) {
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
}
type Callback func (reqeust string)
type OutputData struct {
    Time float64
    Path string`json:"Path"`
    RemoteAddr string
    //ContentLength int64
    Message string
    Table Table

}
type Table struct {
    Columns []string
    Rows [][]interface{}
}
func getOutputData(req *http.Request)OutputData{
    outputData:=OutputData{}
    outputData.Time=float64(time.Now().UnixNano()) / (1000 * 1000 * 1000)
    outputData.Path=req.URL.Path
    outputData.RemoteAddr=req.RemoteAddr
    //outputData.ContentLength=req.ContentLength
    return outputData
}
//请求处理函数
func http_handler_proc(w http.ResponseWriter, req *http.Request) {
    //获取请求资源
    outputData := getOutputData(req)
    req.ParseForm()
    proc := req.FormValue("proc")
    if proc==""{
        outputData.Table = GetParamData()
        //outputData.Message="请提供proc"
        WriteJSON(w, outputData)
        return
    }
    //outputData.Data = make(map[interface{}]interface{})
    //outputData.Data["123"] = 123;
    outputData.Table = GetProcData(proc)
    //fmt.Println(outputData)
    WriteJSON(w, outputData)
}

func http_handler_root(w http.ResponseWriter, req *http.Request) {
    t, err := template.ParseFiles("template/html/index.html")
    checkErr(err)
    t.Execute(w, nil)
}
func WriteJSON(w http.ResponseWriter, outputData OutputData){
    //支持全域名访问，不安全，部署后需要固定限制为客户端网址
    w.Header().Set("Access-Control-Allow-Origin", "*");
    if false{
        //outputData.Body=buf.String()
        w.Header().Set("content-type", "application/json")
    }
    outputJSON,err:=json.Marshal(outputData)
    if err!=nil{
        //fmt.Println(string(outputJSON))
        return
    }
    io.WriteString(w, string(outputJSON))
}
