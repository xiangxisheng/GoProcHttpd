package db

import (
    "strconv"
    "fmt"
    "database/sql"
)

type Table struct {
    Columns []string
    Rows [][]interface{}
}

func GetTableBySql(sqlstr string, params []interface{}) (Table, error) {
    table := Table{}
    stmt, err := sqlDB.Prepare(sqlstr)
    if checkErr(err, "sqlDB.Prepare") { return table, err }
    // Execute the query
    rows, err := stmt.Query(params...)
    if checkErr(err, "stmt.Query") { return table, err }
    TimerReset()
    var table_rows [][]interface{}
    // Get column names
    columns, err := rows.Columns()
    if checkErr(err, "rows.Columns") { return table, err }
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
            if checkErr(err, "rows.Scan") { return table, err }
        }
        // Now do something with the data.
        // Here we just print each column as a string.
        var sValues []interface{}
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            var value interface{}
            switch GetTypeByBytes(col){
                case "nil":
                    value = nil
                    break
                case "num":
                    value, err = strconv.Atoi(string(col))
                    if checkErr(err, "strconv.Atoi") { return table, err }
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
                fmt.Print("\n", columns[i], ": ", value)
            }
        }
        table_rows = append(table_rows, sValues)
    }
    err = rows.Err()
    if checkErr(err, "rows.Err") { return table, err }
    /*
    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }//*/
    rows.Close()
    table.Rows = table_rows
    return table, err
}

