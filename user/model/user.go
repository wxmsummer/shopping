package model

import (
	"github.com/jinzhu/gorm"
)

// 用户结构体
type User struct {
	gorm.Model
	Name     string
	Phone    string `gorm:"type:char(11);`
	Password string
	Points   int32 // 用户积分
	Level    int32 // 用户等级
}
