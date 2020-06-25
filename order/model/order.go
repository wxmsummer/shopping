package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID     int32   // 订单用户id
	ProductID  string // 订单商品id：一个订单可包含多个商品
	CreateTime string  // 订单创建时间
	State      string  // 订单状态：未支付、已支付
}