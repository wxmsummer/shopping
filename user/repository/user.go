package repository

import (
	"github.com/jinzhu/gorm"
	"shopping/user/model"
)

// 用户数据库操作相关接口
type Repository interface {
	FindByID(int32) (*model.User, error)                     // 根据id查找用户
	Create(*model.User) error                                // 创建用户
	Update(*model.User) (*model.User, error)                 // 更新用户
	FindByField(string, string, string) (*model.User, error) // 条件查询
}

type User struct {
	Db *gorm.DB
}

func (repo *User) FindByID(id int32) (*model.User, error) {
	user := &model.User{}
	user.ID = uint(id)
	err := repo.Db.First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *User) Create(user *model.User) error {
	err := repo.Db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *User) Update(user *model.User) (*model.User, error) {
	err := repo.Db.Model(user).Update(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *User) FindByField(key, value, fields string) (*model.User, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	user := &model.User{}
	err := repo.Db.Select(fields).Where(key+" = ?", value).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
