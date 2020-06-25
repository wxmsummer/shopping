package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID          string  // 订单用户id
	ProductID       string  // 订单商品id：一个订单可包含多个商品
	CreateTime      int64   // 订单创建时间
	State           int32   // 订单状态：未支付、已支付、已撤销
	TotalPrice      float32 // 订单总价
	ShippingAddress string  // 收货地址
}

const (
	OrderStateNotPay   = 0 // 未支付
	OrderStateHasPayed = 1 // 已支付
	OrderStateDeal     = 2 // 已成交
	OrderStateCanceled = 3 // 已撤销
)
