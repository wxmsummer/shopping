package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	"net/http"
	"strconv"
	"time"

	proto "github.com/wxmsummer/shopping/order/proto/order"
)

// 获取首页
func GetCreateOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "place_order.html", nil)
}

func PostCreateOrder(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// 从前端请求获取数据，初始化order
	order := &proto.Order{
		Id:         0,
		UserID:     0,
		ProductID:  c.PostForm("product_id"),
		CreateTime: time.Now().Unix(),
		State:      0,
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

func GetOrderById(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// 从前端请求获取orderID
	orderID, _ := strconv.Atoi(c.Query("order_id"))

	// call the backend service
	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.GetOrderById(context.TODO(), &proto.GetOrderByIdReq{
		OrderID: int32(orderID),
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}

	c.JSON(200, rsp)
}

func GetOrdersByProductId(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	userId, _ := strconv.Atoi(c.Query("user_id"))
	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.GetOrdersByUserId(context.TODO(), &proto.GetOrdersByUserIdReq{
		UserID: int32(userId),
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

	orderId, _ := strconv.Atoi(c.Query("order_id"))
	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.CancelOrder(context.TODO(), &proto.CancelOrderReq{
		OrderID: int32(orderId),
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}

	c.JSON(200, rsp)
}