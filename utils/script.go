package utils

import (
    "log"
    "net/http"
    "sync"
    "time"
)

var (
    successRateMu sync.Mutex
    successRate   float64
)

func SendRequests(qps int) {
    var wg sync.WaitGroup
    ticker := time.NewTicker(time.Second / time.Duration(qps))
    defer ticker.Stop()

    totalRequests := 0
    successfulRequests := 0

    for {
        select {
        case <-ticker.C:
            totalRequests++
            wg.Add(1)
            go func() {
                defer wg.Done()
                resp, err := http.Get("http://localhost:8888/students?pageSize=5&prefix=%E8%83%A1")
                if err == nil && resp.StatusCode == http.StatusOK {
                    successfulRequests++
                }
            }()
            
            successRateMu.Lock()
            successRate = float64(successfulRequests) / float64(totalRequests)
            successRateMu.Unlock()
            log.Printf("Success Rate: %.2f%%\n", successRate*100)
        }

    }
}

func MonitorSuccessRate() {
    qps := 10000 // 设定QPS
    SendRequests(qps)
}
