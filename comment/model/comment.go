package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	userID int32    // 评价用户id
	productID int32  // 评价商品id
	content string    // 评价内容
	createTime string // 评价创建时间
}
