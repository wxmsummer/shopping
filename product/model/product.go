package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name         string
	Classify     string
	Tag          string
	Price        string
	SalesVolume  int32
	CommentsNum  int32
	Inventory    int32
	Introduction string
}
