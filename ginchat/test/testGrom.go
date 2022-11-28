package main

import (
	"chat/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		log.Fatalln("数据库连接十失败: error:", err)
	}
	db.AutoMigrate(&model.UserBasic{})
	user := &model.UserBasic{}
	user.Name = "付志东"
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	db.Create(user)
}
