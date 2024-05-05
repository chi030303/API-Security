package controller

import (
	"net/http"
	"strconv"
	"API-Security/models"
	"API-Security/utils"
	"github.com/gin-gonic/gin"
	"errors"
)

func checkQuery(prefix,birth,pageSizeStr string) (int,error) {
	// 检查参数是否为空
    if prefix == "" && birth == "" {
        return 0, errors.New("必须提供姓名前缀或出生年份")
    }

    // 解析 pageSize 参数
    pageSize, err := strconv.Atoi(pageSizeStr)
    if err != nil || pageSize < 1 || pageSize > 1000 {
        return 0, errors.New("pageSize 参数无效")
    }

    return pageSize, nil
}

// 将出参传递给日志输出函数
func PassResultToUtils(result interface{}) {
    utils.LastResult = result
}

func SearchInfo(c *gin.Context) {
	// 获取查询参数
	prefix := c.Query("prefix")
	birthYearStr := c.Query("birth")
	pageSizeStr := c.Query("pageSize")


	// 检查参数是否有效
	pageSize, err := checkQuery(prefix, birthYearStr, pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从数据库中读取学生信息
	students, err := models.QueryStudents(prefix, birthYearStr, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	// 检查是否查询到学生信息
	if len(students) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "未找到匹配的学生信息"})
		return
	}

	// 将结果传递给 utils 包
	PassResultToUtils(students)

	// 返回查询结果
	c.JSON(http.StatusOK, students)
}
