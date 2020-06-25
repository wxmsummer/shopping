package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	proto "github.com/wxmsummer/shopping/user/proto/user"
	"net/http"
	"strconv"
)

// 获取注册界面
func GetUserRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "user_register.html", nil)
}

// 获取登录界面
func GetUserLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "user_login.html", nil)
}

// 获取用户中心界面
func GetUserCenter(c *gin.Context) {
	c.HTML(http.StatusOK, "user_center.html", nil)
}

// 获取用户中心界面
func GetMyInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "user_myInfo.html", nil)
}

// 获取用户中心界面
func GetMyOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "user_myOrder.html", nil)
}

// 获取用户中心界面
func GetMyAddress(c *gin.Context) {
	c.HTML(http.StatusOK, "user_myAddress.html", nil)
}

// 提交表单进行注册
func PostUserRegister(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// call the backend service
	client := proto.NewUserService("go.micro.service.user", server.Client())
	user := proto.User{
		Phone: c.PostForm("phone"),
		Password: c.PostForm("password"),
	}
	rsp, err := client.Register(context.TODO(), &proto.RegisterReq{
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
func PostUserLogin(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// call the backend service
	client := proto.NewUserService("go.micro.service.user", server.Client())

	phone := c.PostForm("phone")
	rsp, err := client.Login(context.TODO(), &proto.LoginReq{
		Phone: phone,
		Password: c.PostForm("password"),
	})
	fmt.Println("xxxxx 1")

	if err != nil {
		c.JSON(500, rsp)
		return
	}
	fmt.Println("xxxxx 2")

	// 登录成功，设置cookie，这里直接将cookie设置为phone（需完善）
	c.SetCookie("user_cookie", phone, 3600, "/","localhost", false, true)

	// 登录成功直接跳转到首页
	c.Redirect(302, "/")
}

// 退出登录
func UserLogout(c *gin.Context) {

	// 登出成功，重置cookie
	c.SetCookie("user_cookie", "", -1, "/","localhost", false, true)

	// 登录成功直接跳转到首页
	c.Redirect(302, "/")

}

func GetUserLevel(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	userID, _ := strconv.Atoi(c.Param("userID"))

	// call the backend service
	client := proto.NewUserService("go.micro.service.user", server.Client())
	rsp, err := client.GetLevel(context.TODO(), &proto.GetLevelReq{
		Id: int32(userID),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	c.JSON(http.StatusOK, rsp)
}