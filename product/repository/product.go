package repository

import (
	"github.com/jinzhu/gorm"
	"shopping/product/model"
)

// 用户数据库操作相关接口
type Repository interface {
	FindByID(int32) (*model.Product, error)                     // 根据id查找商品
	Create(*model.Product) error                                // 创建商品
	Update(*model.Product) (*model.Product, error)              // 更新商品
	FindByField(string, string, string) (*model.Product, error) // 条件查询
}

type Product struct {
	Db *gorm.DB
}

func (repo *Product) FindByID(id string) (*model.Product, error) {
	product := &model.Product{}
	product.ProductID = id
	err := repo.Db.First(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *Product) Create(product *model.Product) error {
	err := repo.Db.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *Product) Update(product *model.Product) (*model.Product, error) {
	err := repo.Db.Model(product).Update(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *Product) FindByField(key, value, fields string) (*model.Product, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	product := &model.Product{}
	err := repo.Db.Select(fields).Where(key+" = ?", value).First(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *Product) Delete(id string) error {

	var product Product
	err := repo.Db.Delete(&product).Where("product_id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}