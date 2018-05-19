package db

import (
    "fmt"
    "database/sql"
)

func GetTypeByBytes(rawBytes sql.RawBytes)string {
    if len(rawBytes) == 0 {
        return "string"
    }
    if rawBytes == nil {
        return "nil"
    }
    var min byte = 255
    var max byte = 0
    for i, char := range rawBytes {
        if false{
            fmt.Println(i, ", ", char)
        }
        if char < min {
            min = char
        }
        if char > max {
            max = char
        }
    }
    if min < 32 {
        return "base64"
    }
    if min < 48 {
        return "string"
    }
    if max > 57 {
        return "string"
    }
    return "num"
}

