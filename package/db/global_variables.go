package db

var globalVariables map[string]string

func GetGlobalVariables() {
    rows, err := sqlDB.Query("SHOW GLOBAL VARIABLES")
    if checkErr(err, "GetGlobalVariables.Query") { return }
    var Variable_name string
    var Value string
    globalVariables = make(map[string]string)
    for rows.Next() {
        err = rows.Scan(&Variable_name, &Value)
        if checkErr(err, "GetGlobalVariables.Scan") { return }
        globalVariables[Variable_name] = Value
    }
}

