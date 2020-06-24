package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	proto "github.com/wxmsummer/shopping/user/proto/user"
	"net/http"
	"strconv"
)

// 获取注册界面
func GetRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "user_register.html", nil)
}

// 获取登录界面
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "user_login.html", nil)
}

// 提交表单进行注册
func PostRegister(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// call the backend service
	webClient := proto.NewUserService("go.micro.service.user", server.Client())
	user := proto.User{
		Phone: c.PostForm("phone"),
		Password: c.PostForm("password"),
	}
	rsp, err := webClient.Register(context.TODO(), &proto.RegisterReq{
		User: &user,
	})
	if err != nil {
		// 应该跳转到错误界面
		c.JSON(500, rsp)
		return
	}

	//c.JSON(http.StatusOK, rsp)

	// 注册成功直接跳转到登录界面
	c.Redirect(302, "/user/login")
}

// 提交表单进行登录
func PostLogin(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// call the backend service
	webClient := proto.NewUserService("go.micro.service.user", server.Client())

	phone := c.PostForm("phone")
	rsp, err := webClient.Login(context.TODO(), &proto.LoginReq{
		Phone: phone,
		Password: c.PostForm("password"),
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}

	// 登录成功，设置cookie，这里直接将cookie设置为phone（需完善）
	c.SetCookie("user_cookie", phone, 3600, "/","localhost", false, true)

	// 登录成功直接跳转到首页
	c.Redirect(302, "/")
}

// 退出登录
func Logout(c *gin.Context) {

	// 登出成功，重置cookie
	c.SetCookie("user_cookie", "", -1, "/","localhost", false, true)

	// 登录成功直接跳转到首页
	c.Redirect(302, "/")

}

func GetLevel(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	userID, _ := strconv.Atoi(c.Param("userID"))

	// call the backend service
	webClient := proto.NewUserService("go.micro.service.user", server.Client())
	rsp, err := webClient.GetLevel(context.TODO(), &proto.GetLevelReq{
		Id: int32(userID),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	c.JSON(http.StatusOK, rsp)
}