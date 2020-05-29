package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/errors"

	"golang.org/x/crypto/bcrypt"
	"shopping/user/model"
	"shopping/user/repository"

	log "github.com/micro/go-micro/v2/logger"

	proto "github.com/wxmsummer/shopping/proto/user"
)

type User struct{ Repo *repository.User }

// 用户注册
func (e *User) Register(ctx context.Context, req *proto.RegisterReq, rsp *proto.Resp) error {

	// 对密码进行hash
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 先从数据库查看手机号对应的账号是否已被注册
	isExistUser, err := e.Repo.FindByField("phone", req.User.Phone, "")
	if err != nil {
		log.Error("e.Repo.FindByField err")
		return err
	}
	// 若能从数据库中查找到账号
	if isExistUser != nil {
		return errors.Unauthorized("go.micro.service.user.register","账号已存在，注册失败！")
	}

	// 初始化user
	user := &model.User{
		Name:     req.User.Name,
		Phone:    req.User.Phone,
		Password: string(hashPwd),
		Points:   0, // 用户积分初始化为0
		Level:    1, // 用户等级初始化为1
	}

	// 注册账号到数据库
	err = e.Repo.Create(user)
	if err != nil {
		log.Error("e.Repo.Create(proto) err")
		return err
	}

	rsp.Code = 200
	rsp.Msg = "注册成功！"
	return nil
}

func (e *User) Login(ctx context.Context, req *proto.LoginReq, rsp *proto.Resp) error {
	user, err := e.Repo.FindByField("phone", req.Phone, "id , password")
	if err != nil {
		return err
	}
	if user == nil {
		return errors.Unauthorized("go.micro.service.user.login", "该账号不存在")
	}
	// 对比密码是否一致
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return errors.Unauthorized("go.micro.service.user.login", "密码错误")
	}

	rsp.Code = 200
	rsp.Msg = "登录成功！"

	return nil
}

func (e *User) Logout(ctx context.Context, req *proto.LogoutReq, rsp *proto.Resp) error {
	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("%d logout success!", req.Id)
	return nil
}

func (e *User) GetLevel(ctx context.Context, req *proto.GetLevelReq, rsp *proto.Resp) error {
	return nil
}