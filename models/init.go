package models

import (
	"fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库连接
func InitDB() (error) {
    db, err := gorm.Open(sqlite.Open("security.db"), &gorm.Config{})
    if err != nil {
        return fmt.Errorf("无法连接数据库：%w", err)
    }
    DB = db
    return nil
}