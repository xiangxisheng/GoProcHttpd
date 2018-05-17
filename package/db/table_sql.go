package db

import (
    "bytes"
)

func GetParamData() Table {
    var buf bytes.Buffer
    buf.WriteString("SELECT T.SPECIFIC_SCHEMA,T.SPECIFIC_NAME,MAX(PARAMETERS)PARAMETERS FROM (")
    buf.WriteString("SELECT t1.SPECIFIC_SCHEMA,t1.SPECIFIC_NAME,GROUP_CONCAT(t1.PARAMETER_NAME,'|',t1.DATA_TYPE,'|',IFNULL(t1.CHARACTER_MAXIMUM_LENGTH,''),IFNULL(t1.NUMERIC_PRECISION,'')ORDER BY ORDINAL_POSITION ASC)PARAMETERS FROM `information_schema`.`PARAMETERS` t1 GROUP BY t1.SPECIFIC_SCHEMA,t1.SPECIFIC_NAME")
    buf.WriteString(" UNION ")
    buf.WriteString("SELECT t2.ROUTINE_SCHEMA,t2.ROUTINE_NAME,NULL FROM `information_schema`.`ROUTINES` t2 WHERE t2.ROUTINE_TYPE='PROCEDURE'")
    buf.WriteString(")T GROUP BY T.SPECIFIC_SCHEMA,T.SPECIFIC_NAME ORDER BY T.SPECIFIC_NAME ASC")
    return GetTableBySql(buf.String())
}

func GetProcData(procName string) Table {
    sqlstr := "CALL `" + procName + "`()"
    return GetTableBySql(sqlstr)
}

