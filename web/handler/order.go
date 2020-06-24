package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	"net/http"
	"time"

	proto "github.com/wxmsummer/shopping/order/proto/order"
)

// 获取首页
func GetPlaceOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "place_order.html", nil)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()
	
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求获取数据，初始化order
	order := &proto.Order{
		Id:         0,
		UserID:     0,
		ProductID:  nil,
		CreateTime: "",
		State:      "",
	}

	// call the backend service
	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.CreateOrder(context.TODO(), &proto.CreateOrderReq{
		Order: order,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()
	
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求获取orderID
	orderID := request["orderID"].(int32)

	// call the backend service
	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.GetOrderById(context.TODO(), &proto.GetOrderByIdReq{
		OrderID: orderID,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()
	
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	userID := request["userID"].(int32)

	// call the backend service
	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.GetAllOrders(context.TODO(), &proto.GetAllOrdersReq{
		UserID: userID,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()
	
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求获取orderID
	orderID := request["orderID"].(int32)

	// call the backend service
	webClient := proto.NewOrderService("go.micro.service.order", server.Client())
	rsp, err := webClient.CancelOrder(context.TODO(), &proto.CancelOrderReq{
		OrderID: orderID,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}