package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/errors"
	"golang.org/x/crypto/bcrypt"
	"shopping/user/model"
	"shopping/user/repository"

	log "github.com/micro/go-micro/v2/logger"

	proto "shopping/user/proto/user"
)

type User struct{ Repo *repository.User }

// 用户注册
func (e *User) Register(ctx context.Context, req *proto.RegisterReq, rsp *proto.Resp) error {

	// 先从数据库查看手机号对应的账号是否已被注册
	isExistUser, err := e.Repo.FindByField("phone", req.User.Phone, "")
	// 若能从数据库中查找到账号
	if isExistUser != nil {
		rsp.Code = 501
		rsp.Msg = "账号已存在，注册失败！"
	} else {
		// 初始化user
		user := &model.User{
			UserId: req.User.UserId,
			Name:     req.User.Name,
			Phone:    req.User.Phone,
			Password: req.User.Password,
			Points:   0, // 用户积分初始化为0
			Level:    1, // 用户等级初始化为1
		}

		// 注册账号到数据库
		err = e.Repo.Create(user)
		if err != nil {
			log.Error("e.Repo.Create(proto) err")
			return err
		}
	}

	rsp.Code = 200
	rsp.Msg = "注册成功！"
	return nil
}

func (e *User) Login(ctx context.Context, req *proto.LoginReq, rsp *proto.LoginResp) error {
	user, err := e.Repo.FindByField("phone", req.Phone, "password")
	if err != nil {
		return err
	}
	if user == nil {
		rsp.Code = 500
		rsp.Msg = "该账号不存在"
		return errors.Unauthorized("go.micro.service.user.login", "该账号不存在")
	}
	// 对比密码是否一致
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		rsp.Code = 500
		rsp.Msg = "密码错误"
		return errors.Unauthorized("go.micro.service.user.login", "密码错误")
	}

	rsp.Code = 200
	rsp.Msg = "登录成功！"
	rsp.User = &proto.User{
		UserId:   user.UserId,
		Name:     user.Name,
		Phone:    user.Phone,
		Password: "",
		Points:   user.Points,
		Level:    user.Level,
	}

	return nil
}

func (e *User) GetUserByPhone(ctx context.Context, req *proto.GetUserByPhoneReq, rsp *proto.GetUserByPhoneResp) error {
	user, err := e.Repo.FindByField("phone", req.Phone, "")
	if err != nil {
		return err
	}
	if user == nil {
		rsp.Code = 500
		rsp.Msg = "该账号不存在"
		return errors.Unauthorized("go.micro.service.user.login", "该账号不存在")
	}

	rsp.Code = 200
	rsp.Msg = "登录成功！"

	return nil
}

func (e *User) Logout(ctx context.Context, req *proto.LogoutReq, rsp *proto.Resp) error {
	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("%s logout success!", req.UserId)
	return nil
}

func (e *User) GetLevel(ctx context.Context, req *proto.GetLevelReq, rsp *proto.Resp) error {
	return nil
}