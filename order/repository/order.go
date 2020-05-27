package repository

import (
	"github.com/jinzhu/gorm"
	"shopping/order/model"
)

// 用户数据库操作相关接口
type Repository interface {
	FindByID(int32) (*model.Order, error)                     	// 根据id查找订单
	Create(*model.Order) error                                	// 创建订单
	Update(*model.Order) (*model.Order, error)              	// 更新订单：撤销订单即将订单状态修改为已撤销
	FindByField(string, string, string) (*model.Order, error) 	// 条件查询
}

type Order struct {
	Db *gorm.DB
}

func (repo *Order) FindByID(id int32) (*model.Order, error) {
	order := &model.Order{}
	order.ID = uint(id)
	err := repo.Db.First(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *Order) Create(order *model.Order) error {
	err := repo.Db.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *Order) Update(order *model.Order) (*model.Order, error) {
	err := repo.Db.Model(order).Update(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *Order) FindByField(key, value, fields string) (*model.Order, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	order := &model.Order{}
	err := repo.Db.Select(fields).Where(key+" = ?", value).First(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *Order) Delete(id int32) error {

	var order Order
	err := repo.Db.Delete(&order).Where("id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}