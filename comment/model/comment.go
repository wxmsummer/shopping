package model

import (
	"github.com/jinzhu/gorm"
)

// 一个订单对应一个评论
// 一个商品对应多个订单
type Comment struct {
	gorm.Model
	UserID     string // 评价用户id
	ProductID  string // 评价商品id
	OrderID    string // 该商品的订单id
	Star       int32  // 评价星级：1-5
	Content    string // 评价内容
	CreateTime int64  // 评价创建时间
}
