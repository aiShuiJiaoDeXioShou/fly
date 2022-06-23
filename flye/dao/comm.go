package dao

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 如果没有特殊原因，一般情况下comm都是fly项目包的初始化模块
// 定义一个全局DB对象，连接数据库
var (
	DB        *gorm.DB
	err       error
	newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: false,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)
)

func init() {
	DB, err = gorm.Open(sqlite.Open("FLY.DB"), &gorm.Config{
		// 是否启用Logger
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		}, // 使用单数表名
	})
	if err != nil {
		panic(err)
	}
}
