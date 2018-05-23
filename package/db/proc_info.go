package db

import (
    //"fmt"
    "strings"
    "database/sql"
    "strconv"
)

//ProcParam
type ProcParam struct {
    PARAMETER_NAME string
    DATA_TYPE string
    MAXIMUM_LENGTH int
}

func isUnderline(r rune) bool {
    return r == '_'
}

//取得【存储过程】要求传递的参数
func GetProcParam(sDbName string, sProcName string) ([]ProcParam, error) {
    var aProcParam []ProcParam
    sqlstr := "SELECT PARAMETER_NAME,DATA_TYPE,CHARACTER_MAXIMUM_LENGTH,NUMERIC_PRECISION,NUMERIC_SCALE FROM `information_schema`.`PARAMETERS` WHERE SPECIFIC_NAME=? AND SPECIFIC_SCHEMA=?"
    // 1：预编译
    stmt, err := sqlDB.Prepare(sqlstr)
    if checkErr(err, "GetProcParam.sqlDB.Prepare") { return aProcParam, err }
    // 2：执行查询
    rows, err := stmt.Query(sProcName, sDbName)
    if checkErr(err, "GetProcParam.stmt.Query") { return aProcParam, err }
    // 3：取得列
    columns, err := rows.Columns()
    if checkErr(err, "GetProcParam.rows.Columns") { return aProcParam, err }
    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }
    for rows.Next() {
        oProcParam := ProcParam{}
        // 1：扫描每一行数据
        err := rows.Scan(scanArgs...);
        if checkErr(err, "GetProcParam.rows.Scan") { return aProcParam, err }
        // 2：取得每一行数据
        oProcParam.PARAMETER_NAME = strings.TrimLeftFunc(string(values[0]), isUnderline)
        oProcParam.DATA_TYPE = string(values[1])
        if values[2] != nil {
            oProcParam.MAXIMUM_LENGTH, err = strconv.Atoi(string(values[2]))
            if checkErr(err, "GetProcParam.values[2].strconv") { return aProcParam, err }
        } else if values[3] != nil {
            oProcParam.MAXIMUM_LENGTH, err = strconv.Atoi(string(values[3]))
            if checkErr(err, "GetProcParam.values[3].strconv") { return aProcParam, err }
            if values[4] != nil {
                // 把小数部分也计算进去
                i, err := strconv.Atoi(string(values[4]))
                if checkErr(err, "GetProcParam.values[4].strconv") { return aProcParam, err }
                if i > 0 {
                    // 这里包含小数点的长度
                    oProcParam.MAXIMUM_LENGTH += i + 1
                }
            }
        }
        // 3：追加到数组
        aProcParam = append(aProcParam, oProcParam)
    }
    return aProcParam, err
}

