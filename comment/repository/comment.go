package repository

import (
	"github.com/jinzhu/gorm"
	"shopping/comment/model"
)

// 用户数据库操作相关接口
type Repository interface {
	FindByID(int32) (*model.Comment, error)                     // 根据id查找商品评价
	Create(*model.Comment) error                                // 创建商品评价
	FindByField(string, string, string) (*model.Comment, error) // 条件查询评价
}

type Comment struct {
	Db *gorm.DB
}

func (repo *Comment) FindByID(id int32) (*model.Comment, error) {
	comment := &model.Comment{}
	comment.ID = uint(id)
	err := repo.Db.First(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (repo *Comment) Create(comment *model.Comment) error {
	err := repo.Db.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}


func (repo *Comment) FindByField(key, value, fields string) (*model.Comment, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	comment := &model.Comment{}
	err := repo.Db.Select(fields).Where(key+" = ?", value).First(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}
