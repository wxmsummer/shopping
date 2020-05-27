package handler

import (
	"context"
	"github.com/micro/go-micro/v2/errors"
	"golang.org/x/crypto/bcrypt"
	"shopping/user/model"
	"shopping/user/repository"

	log "github.com/micro/go-micro/v2/logger"

	user "shopping/user/proto/user"
)

type User struct{ Repo *repository.User }

// 用户注册
func (e *User) Register(ctx context.Context, req *user.RegisterReq, rsp *user.Resp) error {

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     req.User.Name,
		Phone:    req.User.Phone,
		Password: string(hashPwd),
		Points:   0, // 用户积分初始化为0
		Level:    1, // 用户等级初始化为1
	}

	err = e.Repo.Create(user)
	if err != nil {
		log.Error("e.Repo.Create(user) err")
		return err
	}

	rsp.Code = 200
	rsp.Msg = "注册成功！"
	return nil
}

func (e *User) Login(ctx context.Context, req *user.LoginReq, rsp *user.Resp) error {
	user, err := e.Repo.FindByField("phone", req.Phone, "id , password")
	if err != nil {
		return err
	}
	if user == nil {
		return errors.Unauthorized("go.micro.srv.user.login", "该账号不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return errors.Unauthorized("go.micro.srv.user.login", "密码错误")
	}

	rsp.Code = 200
	rsp.Msg = "登录成功！"

	return nil
}

func (e *User) Logout(ctx context.Context, req *user.LogoutReq, rsp *user.Resp) error {
	return nil
}

func (e *User) GetLevel(ctx context.Context, req *user.GetLevelReq, rsp *user.Resp) error {
	return nil
}