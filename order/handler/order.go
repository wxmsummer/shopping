package handler

import (
	"context"
	"shopping/order/repository"

	proto "shopping/order/proto/order"
)

type Order struct{ Repo *repository.Order }

func (e *Order) CreateOrder(ctx context.Context, req *proto.CreateOrderReq, rsp *proto.Resp) error {

	rsp.Code = 200
	rsp.Msg = "CreateOrder成功！"
	return nil
}

func (e *Order) GetOrderById(ctx context.Context, req *proto.GetOrderByIdReq, rsp *proto.GetOrderByIdResp) error {

	rsp.Code = 200
	rsp.Msg = "GetOrderById成功！"
	return nil
}

func (e *Order) GetAllOrders(ctx context.Context, req *proto.GetAllOrdersReq, rsp *proto.GetAllOrdersResp) error {

	rsp.Code = 200
	rsp.Msg = "GetAllOrders成功！"
	return nil
}

func (e *Order) CancelOrder(ctx context.Context, req *proto.CancelOrderReq, rsp *proto.Resp) error {

	rsp.Code = 200
	rsp.Msg = "CancelOrder成功！"
	return nil
}
