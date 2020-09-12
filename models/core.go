package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var (
	db *gorm.DB
)

func init() {
	//logs.Info("----------init 方法 ------------")
	var err error
	err = os.MkdirAll("data", 0777)
	if err!=nil {
		panic("失败" + err.Error())
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("连接数据库失败")
	}

	//同步表结构
	db.AutoMigrate(&User{})
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Name: "小孙老师",
			Email: "admin@qq.com",
			Pwd: "123123",
			Avatar: "/static/images/info_img.png",
			Role: 0,
		})
	}
}

