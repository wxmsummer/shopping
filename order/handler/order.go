package handler

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	"shopping/order/model"
	"shopping/order/repository"

	proto "shopping/order/proto/order"
)

type Order struct{ Repo *repository.Order }

func (e *Order) CreateOrder(ctx context.Context, req *proto.CreateOrderReq, rsp *proto.Resp) error {


	order := model.Order{
		UserID:     	req.Order.UserID,
		ProductID:  	req.Order.ProductID,
		CreateTime:		req.Order.CreateTime,
		State:      	req.Order.State,
		TotalPrice:      req.Order.TotalPrice,
		ShippingAddress: req.Order.ShippingAddress,
	}

	err := e.Repo.Create(&order)
	if err != nil {
		log.Error("e.Repo.Create(&order) err")
		return err
	}
	rsp.Code = 200
	rsp.Msg = "CreateOrder success！"
	return nil

}

func (e *Order) GetOrderById(ctx context.Context, req *proto.GetOrderByIdReq, rsp *proto.GetOrderByIdResp) error {

	var order proto.Order
	err := e.Repo.Db.Where("order_id = ?" , req.OrderID).Find(&order).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = "GetOrderById success"
	rsp.Order = &order

	return nil
}

func (e *Order) GetOrdersByUserId(ctx context.Context, req *proto.GetOrdersByUserIdReq, rsp *proto.GetOrdersByUserIdResp) error {

	var orders []*proto.Order

	err := e.Repo.Db.Where("user_id = ?" , req.UserID).Find(&orders).Limit(10).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = "GetOrdersByUserId success！"
	rsp.Orders = orders
	return nil
}

func (e *Order) CancelOrder(ctx context.Context, req *proto.CancelOrderReq, rsp *proto.Resp) error {

	var order proto.Order

	err := e.Repo.Db.Model(&order).Where("order_id = ?" , req.OrderID).Update("state", model.OrderStateCanceled).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = "CancelOrder success！"
	return nil
}
