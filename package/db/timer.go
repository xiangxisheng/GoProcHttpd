package db

import (
    "fmt"
    "time"
    "strconv"
)

var timer1 *time.Timer
var timediff1 time.Duration

func Timer() {
    timer1 = time.NewTimer(0)
    for {
        select {
        case <-timer1.C:
            // 数据库每隔一段时间要进行一次操作，防止掉线或被踢
            fmt.Print("\nsqlDB.Ping().")
            timenow1 := time.Now()
            sqlDB.Ping()
            timediff1 = time.Now().Sub(timenow1)
            fmt.Print(".OK! (", (timediff1), ")")
            TimerReset()
        }
    }
}

func TimerReset() {
    timeout, err := strconv.Atoi(globalVariables["wait_timeout"])
    if checkErr(err, "TimerReset.Atoi") { return }
    duration := time.Duration(float64(time.Second) * (float64(timeout) * 0.9))
    timer1.Reset(duration - timediff1)
}

