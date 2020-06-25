package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/rs/xid"
	productProto "github.com/wxmsummer/shopping/product/proto/product"
	userProto "github.com/wxmsummer/shopping/user/proto/user"
	"net/http"
	"strconv"
)

// 获取管理员首页
func GetAdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_index.html", nil)
}

// 获取管理员登录
func GetAdminLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login.html", nil)
}

// 获取管理员添加商品界面
func GetAdminAddProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_addProduct.html", nil)
}

// 获取管理员列出商品界面
func GetAdminListProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_listProduct.html", nil)
}

// 获取管理员更新商品界面
func GetAdminUpdateProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_addProduct.html", nil)
}

// 提交表单进行登录
func PostAdminLogin(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// call the backend service
	client := userProto.NewUserService("go.micro.service.user", server.Client())

	phone := c.PostForm("phone")
	rsp, err := client.Login(context.TODO(), &userProto.LoginReq{
		Phone:    phone,
		Password: c.PostForm("password"),
	})

	if err != nil {
		c.JSON(500, rsp)
		return
	}

	// 登录成功，设置cookie，这里直接将cookie设置为phone（需完善）
	c.SetCookie("user_cookie", phone, 3600, "/", "localhost", false, true)

	// 登录成功直接跳转到首页
	c.Redirect(302, "/")
}

// 退出登录
func AdminLogout(c *gin.Context) {

	// 登出成功，重置cookie
	c.SetCookie("admin_cookie", "", -1, "/", "localhost", false, true)

	// 登录成功直接跳转到首页
	c.Redirect(302, "/")

}

func PostAdminAddProduct(c *gin.Context) {
	// 验证管理员身份
	_, err := c.Cookie("admin_cookie")
	if err != nil {
		c.Redirect(302, "/admin/login")
	}

	server := grpc.NewService()
	server.Init()

	// call the backend service
	client := productProto.NewProductService("go.micro.service.product", server.Client())

	inventory, _ := strconv.Atoi(c.Param("product_inventory"))
	product := productProto.Product{
		ProductID:   xid.New().String(),
		Name:        c.PostForm("product_name"),
		Classify:    c.PostForm("product_classify"),
		Tag:         c.PostForm("product_tag"),
		Price:       c.PostForm("product_price"),
		SalesVolume: 0,
		CommentsNum: 0,
		Inventory:   int32(inventory),
		Describe:    c.PostForm("product_describe"),
	}
	rsp, err := client.AddProduct(context.TODO(), &productProto.AddProductReq{
		Product: &product,
	})
	if err != nil {
		// 应该跳转到错误界面
		c.JSON(500, rsp)
		return
	}

	c.JSON(http.StatusOK, rsp)

	// c.Redirect(302, "/admin/listProduct")

}

func PostAdminUpdateProduct(c *gin.Context) {
	// 验证管理员身份
	_, err := c.Cookie("admin_cookie")
	if err != nil {
		c.Redirect(302, "/admin/login")
	}

	server := grpc.NewService()
	server.Init()

	// call the backend service
	client := productProto.NewProductService("go.micro.service.product", server.Client())

	inventory, _ := strconv.Atoi(c.Param("product_inventory"))
	product := productProto.Product{
		ProductID:   c.PostForm("product_id"),   // 如何获取productID？？
		Name:        c.PostForm("product_name"),
		Classify:    c.PostForm("product_classify"),
		Tag:         c.PostForm("product_tag"),
		Price:       c.PostForm("product_price"),
		SalesVolume: 0,
		CommentsNum: 0,
		Inventory:   int32(inventory),
		Describe:    c.PostForm("product_describe"),
	}
	rsp, err := client.UpdateProduct(context.TODO(), &productProto.UpdateProductReq{
		Product: &product,
	})
	if err != nil {
		// 应该跳转到错误界面
		c.JSON(500, rsp)
		return
	}

	c.JSON(http.StatusOK, rsp)

	//c.Redirect(302, "/admin/listProduct")

}
