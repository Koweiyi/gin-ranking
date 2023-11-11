package dao

import (
	// "database/sql"
	// "time"
	// "example.com/gin-ranking/config"
	logger "go-ranking/pkg"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error

// db  *sql.DB
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ranking?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	if DB.Error != nil {
		logger.Error(map[string]interface{}{"datebase error": DB.Error})
	}

	sqldb, err := DB.DB()
	if err != nil {
		logger.Error(map[string]interface{}{"datebase error": DB.Error})
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqldb.SetMaxIdleConns(10)
	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqldb.SetMaxOpenConns(100)
	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqldb.SetConnMaxLifetime(time.Hour)
}
