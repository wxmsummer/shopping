package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name         string	// 商品名
	Classify     string	// 分类
	Tag          string // 标签
	Price        string // 价格
	SalesVolume  int32	// 销量
	CommentsNum  int32	// 评价数量
	Inventory    int32	// 库存
	Introduction string // 商品简介
}
