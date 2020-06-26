package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	"net/http"
	"strconv"
	"time"
	"github.com/rs/xid"

	proto "github.com/wxmsummer/shopping/order/proto/order"
)

// 获取首页
func GetCreateOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "place_order.html", nil)
}

// 下单
func PostCreateOrder(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	userId, err := c.Cookie("user_id")
	if err != nil {
		c.Redirect(302, "/user/login")
	}

	totalPrice, _ := strconv.ParseFloat(c.PostForm("total_price"), 10)

	// 从前端请求获取数据，初始化order
	order := &proto.Order{
		OrderID:         xid.New().String(),
		UserID:          userId,
		ProductID:       c.PostForm("product_id"),  // 如何获取productID？？前端获取
		ProductNum:		 c.PostForm("product_num"),
		CreateTime:      time.Now().Unix(),
		State:           0,
		TotalPrice:      float32(totalPrice),
		ShippingAddress: c.PostForm("shipping_address"),
	}

	// call the backend service
	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.CreateOrder(context.TODO(), &proto.CreateOrderReq{
		Order: order,
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}

	c.JSON(200, rsp)
}

func GetOrderByOrderId(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// call the backend service
	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.GetOrderByOrderId(context.TODO(), &proto.GetOrderByOrderIdReq{
		OrderID: c.Query("order_id"),
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}

	c.JSON(200, rsp)
}

func GetOrdersByUserId(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.GetOrdersByUserId(context.TODO(), &proto.GetOrdersByUserIdReq{
		UserID: c.Query("user_id"),
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}

	c.JSON(200, rsp)
}

func CancelOrder(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.CancelOrder(context.TODO(), &proto.CancelOrderReq{
		OrderID: c.Query("order_id"),
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}

	c.JSON(200, rsp)
}