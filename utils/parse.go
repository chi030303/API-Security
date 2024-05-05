package utils

import (
    "fmt"
    "strings"
    "time"
)

// 分割输入的时间范围
func ParseDateRange(dateRangeStr string) (string, string, error) {
    // 按空格分隔字符串
    parts := strings.Fields(dateRangeStr)

    // 检查分隔后的部分数量是否为2
    if len(parts) != 2 {
        return "", "", fmt.Errorf("Invalid date range format: %s", dateRangeStr)
    }

    // 解析起始日期
    startDate, err := time.Parse("2006-01-02", parts[0])
    if err != nil {
        return "", "", fmt.Errorf("Error parsing start date: %s", err)
    }

    // 解析结束日期
    endDate, err := time.Parse("2006-01-02", parts[1])
    if err != nil {
        return "", "", fmt.Errorf("Error parsing end date: %s", err)
    }

    // 检查起始日期是否在结束日期之后
    if startDate.After(endDate) {
        return "", "", fmt.Errorf("Invalid date range: start date is after end date")
    }

    // 格式化日期为字符串形式
    startDateStr := startDate.Format("2006-01-02")
    endDateStr := endDate.Format("2006-01-02")

    return startDateStr, endDateStr, nil
}
