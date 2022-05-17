package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	dm8 "gitee.com/jun3372/driver/dm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	dsn := "dm://SYSDBA:SYSDBA@localhost:5236"
	db, err := gorm.Open(dm8.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	type Result struct {
		ID       uint64 `gorm:"ID"`
		USERNAME string `gorm:"USERNAME"`
		PASSWORD string `gorm:"PASSWORD"`
	}

	type test struct {
		Id    uint64 `json:"id"`
		Title string `json:"title"`
		Info  string `json:"info"`
	}

	if err = db.AutoMigrate(&Result{}); err != nil {
		panic(err)
	}

	r1 := Result{USERNAME: "渣渣辉"}
	if err := db.Where(r1).FirstOrCreate(&r1).Error; err != nil {
		panic(err)
	}

	return

	// 	sqlTr := `CREATE TABLE tests(
	//     id int NOT NULL PRIMARY KEY IDENTITY(1, 1),
	//     title VARCHAR(255),
	//     info VARCHAR(255)
	// );`
	// 	if err := db.Exec(sqlTr).Error; err != nil {
	// 		panic(err)
	// 	}

	t1 := test{Title: "测试标题"}
	fmt.Println(db.Migrator().HasTable(t1))
	if err := db.FirstOrCreate(&t1).Error; err != nil {
		panic(err)
	}

	// if err := db.Delete(&t1).Error; err != nil {
	// 	panic(err)
	// }

	var count int64
	var result map[string]interface{}
	db.Table("all_constraints").Where("OWNER = ? and TABLE_NAME = ?", "SYSDBA", "tests").Scan(&result)

	// var count int64
	// if err := db.Model(t1).Count(&count).Error; err != nil {
	// 	panic(err)
	// }

	v, _ := json.Marshal(result)
	fmt.Println("test", t1, count, string(v))
}
