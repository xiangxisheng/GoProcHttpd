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

func GetTableBySql(sqlstr string) Table {
    // Execute the query
    rows, err := sqlDB.Query(sqlstr)
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
            switch GetTypeByBytes(col){
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

