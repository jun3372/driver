package tests

import (
	"log"
	"os"
	"time"

	dm8 "gitee.com/jun3372/driver/dm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MakeDB(dsn string) (*gorm.DB, error) {
	var (
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // 日志级别
			},
		)
	)

	return gorm.Open(dm8.Open(dsn), &gorm.Config{Logger: newLogger})
}
