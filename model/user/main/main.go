package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"ys_go/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3307)/ys_go?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢sql 阈值
			LogLevel:      logger.Info,
			Colorful:      true, // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{})

	if err != nil {
		fmt.Println("err", err)
	}
}
