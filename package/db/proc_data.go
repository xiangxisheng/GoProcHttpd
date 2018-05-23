package db

import (
    "bytes"
    //"fmt"
    "strings"
)

//获取【存储过程】的名称、所属库名、所需参数等信息
func GetParamData() (Table, error) {
    var buf bytes.Buffer
    buf.WriteString("SELECT T.SPECIFIC_SCHEMA,T.SPECIFIC_NAME,MAX(PARAMETERS)PARAMETERS FROM (")
    buf.WriteString("SELECT t1.SPECIFIC_SCHEMA,t1.SPECIFIC_NAME,GROUP_CONCAT(t1.PARAMETER_NAME,'|',t1.DATA_TYPE,'|',IFNULL(t1.CHARACTER_MAXIMUM_LENGTH,''),IFNULL(t1.NUMERIC_PRECISION,'')ORDER BY ORDINAL_POSITION ASC)PARAMETERS FROM `information_schema`.`PARAMETERS` t1 GROUP BY t1.SPECIFIC_SCHEMA,t1.SPECIFIC_NAME")
    buf.WriteString(" UNION ")
    buf.WriteString("SELECT t2.ROUTINE_SCHEMA,t2.ROUTINE_NAME,'' FROM `information_schema`.`ROUTINES` t2 WHERE t2.ROUTINE_TYPE='PROCEDURE'")
    buf.WriteString(")T GROUP BY T.SPECIFIC_SCHEMA,T.SPECIFIC_NAME ORDER BY T.SPECIFIC_NAME ASC")
    //fmt.Println(buf.String())
    var params []interface{}
    return GetTableBySql(buf.String(), params)
}

//执行【存储过程】并获取数据
func GetProcData(sDbName string, sProcName string, params []interface{}) (Table, error) {
    aSign := make([]string, len(params))
    for i := range aSign {
        aSign[i] = "?"
    }
    sSign := strings.Join(aSign, ",")
    sqlstr := "CALL `" + sDbName + "`.`" + sProcName + "`(" + sSign + ")"
    return GetTableBySql(sqlstr, params)
}

