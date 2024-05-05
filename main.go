package main

import (
    "log"

    "API-Security/models"
    "API-Security/router"
    "API-Security/utils"
)

func main() {
    // 初始化数据库连接
    err := models.InitDB()
    if err != nil {
        log.Fatalf("无法初始化数据库连接：%v", err)
    }

    // 初始化路由
    r := router.Router()

    // 启动请求发送和成功率监控
    go utils.MonitorSuccessRate()

    // 启动一个 goroutine，定时输出每秒的请求数
    go utils.MonitorRequestsPerSecond()

    // 运行路由
    if err := r.Run(":8888"); err != nil {
        log.Panicln(err)
    }
}