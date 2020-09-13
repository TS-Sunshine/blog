package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	// `gorm:"primary_key"` 意思是唯一索引
	Name string `gorm:"primary_key"`
	Pwd string
	Email string `gorm:"primary_key"`
	Avatar string
	// 0代表管理员 1 代表普通用户
	Role int  `gorm:"default:1"`
}

func QueryUserByEmailAndPwd(email, pwd string) (user User, err error) {
	return user, db.Where("email = ? and pwd = ?", email, pwd).Take(&user).Error
}