package utils

import (
	"fmt"
	"sync"
	"time"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 请求计数器相关变量和函数
var (
    requestCountPerSecond int
    lastResetTime         time.Time
    mutex                 sync.Mutex
)

// 初始化
func init() {
    // 启动定时器，每秒重置计数器
    go func() {
        ticker := time.NewTicker(time.Second)
        defer ticker.Stop()
        for {
            <-ticker.C
            // 每秒重置总请求数计数器

            // 每秒重置每秒请求数计数器
            mutex.Lock()
            requestCountPerSecond = 0
            mutex.Unlock()
            lastResetTime = time.Now()
        }
    }()
}

// 定义中间件函数
func RequestCounterMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        mutex.Lock()
        // 在处理请求前递增请求数量
        requestCountPerSecond++
        count := requestCountPerSecond
        mutex.Unlock()

        // 如果请求数量超过 10，则直接返回
        if count > 10 {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
            c.Abort()
            return
        }

        // 继续处理请求
        c.Next()
        fmt.Printf("请求数量: %d\n", count)
    }
}

// 定义函数获取每秒请求数
func GetRequestsPerSecond() int {
    mutex.Lock()
    // 获取当前时间和上次重置时间之间的秒数差
    seconds := int(time.Since(lastResetTime).Seconds())
    count := requestCountPerSecond
    mutex.Unlock()

    // 如果秒数差小于1秒，返回当前秒内的请求数
    if seconds <= 1 {
        return count
    }
    // 如果秒数差大于1秒，返回0
    return 0
}

// 计算每个请求的时间
func ResponseTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求处理开始时间
		startTime := time.Now()

		// 继续处理请求
		c.Next()
		// 计算请求处理时间
		duration := time.Since(startTime)
		// 将处理时间添加到响应头中
		c.Writer.Header().Set("X-Response-Time", fmt.Sprintf("%v", duration))
		fmt.Printf("请求时间: %v\n", duration)
	}
}

// RequestLogInfo 定义请求日志信息结构体
type RequestLogInfo struct {
	LogID       string    `json:"log_id"`
	Method      string    `json:"method"`
	URL         string    `json:"url"`
	Request     string    `json:"request"`
	Response    string    `json:"response"`
	ElapsedTime float64   `json:"elapsed_time"`
	Timestamp   time.Time `json:"timestamp"`
}

// 记录出参
var LastResult interface{}

// LogMiddleware 定义日志中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成唯一的 logID
		logID := uuid.New().String()

		// 记录请求开始时间
		startTime := time.Now()

		// 继续处理请求
		c.Next()

		// 计算请求处理时间
		elapsedTime := time.Since(startTime).Seconds()

		// 构造请求日志信息
		requestInfo := RequestLogInfo{
			LogID:       logID,
			Method:      c.Request.Method,
			URL:         c.Request.URL.String(),
			Request:     fmt.Sprintf("%v", c.Request),
			Response:    fmt.Sprintf("%v", c.Writer.Status()),
			ElapsedTime: elapsedTime,
			Timestamp:   time.Now(),
		}

		// 输出请求日志信息
		fmt.Printf("Request Log Info:\n")
		fmt.Printf("Log ID: %s\n", requestInfo.LogID)
		fmt.Printf("Method: %s\n", requestInfo.Method)
		fmt.Printf("URL: %s\n", requestInfo.URL)
		fmt.Printf("Request: %s\n", requestInfo.Request)
		fmt.Printf("Response: %s\n", requestInfo.Response)
		fmt.Printf("Elapsed Time: %.6f seconds\n", requestInfo.ElapsedTime)
		fmt.Printf("Timestamp: %s\n", requestInfo.Timestamp.Format(time.RFC3339))
		fmt.Printf("output params: %v\n", LastResult)
		fmt.Println()
	}
}

// 每秒输出请求计数
func MonitorRequestsPerSecond() {
    go func() {
        for {
            // 每秒输出一次请求数
            time.Sleep(time.Second)
            requestsPerSecond := GetRequestsPerSecond()
            log.Printf("每秒请求数: %d\n", requestsPerSecond)
        }
    }()
}