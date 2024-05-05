package models

import (
	"fmt"
	"API-Security/utils"
)

// 定义学生结构体
type Student struct {
	ID        uint
	StudentNo int
	Name      string
	Gender    string
	Birth     string
}

// 查询学生信息
func QueryStudents(prefix, birthYearStr string, pageSize int) ([]Student, error) {
    var students []Student
    var startDate, endDate string
    var err error

    // 解析输入的日期
    if birthYearStr != "" {
        startDate, endDate, err = utils.ParseDateRange(birthYearStr)
        fmt.Printf("开始日期: %v, 结束日期: %v\n", startDate, endDate)
        if err != nil {
            return nil, fmt.Errorf("日期输入错误：%s", err)
        }
    }

    // 构建查询条件
    var query string
    var args []interface{}

    // 如果 prefix 不为空，将其作为查询条件
    if prefix != "" {
        query += "name LIKE ?"
        args = append(args, fmt.Sprintf("%s%%", prefix))
    }

    // 如果 birthYearStr 不为空，将其解析的日期范围作为查询条件
    if startDate != "" && endDate != "" {
        if prefix != "" {
            query += " AND " // 如果同时有 prefix 和 date 查询条件，则使用 AND 连接
        }
        query += "(birth BETWEEN DATE(?) AND DATE(?))"
        args = append(args, startDate, endDate)
    }

    // 执行查询操作
    if err := DB.Where(query, args...).Limit(pageSize).Find(&students).Error; err != nil {
        return nil, err
    }

    return students, nil
}
