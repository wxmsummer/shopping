// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order/order.proto

package go_micro_service_order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for OrderService service

type OrderService interface {
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...client.CallOption) (*Resp, error)
	GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...client.CallOption) (*GetOrderByIdResp, error)
	GetOrdersByUserId(ctx context.Context, in *GetOrdersByUserIdReq, opts ...client.CallOption) (*GetOrdersByUserIdResp, error)
	CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...client.CallOption) (*Resp, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.service.order"
	}
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "OrderService.CreateOrder", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...client.CallOption) (*GetOrderByIdResp, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetOrderById", in)
	out := new(GetOrderByIdResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetOrdersByUserId(ctx context.Context, in *GetOrdersByUserIdReq, opts ...client.CallOption) (*GetOrdersByUserIdResp, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetOrdersByUserId", in)
	out := new(GetOrdersByUserIdResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "OrderService.CancelOrder", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	CreateOrder(context.Context, *CreateOrderReq, *Resp) error
	GetOrderById(context.Context, *GetOrderByIdReq, *GetOrderByIdResp) error
	GetOrdersByUserId(context.Context, *GetOrdersByUserIdReq, *GetOrdersByUserIdResp) error
	CancelOrder(context.Context, *CancelOrderReq, *Resp) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		CreateOrder(ctx context.Context, in *CreateOrderReq, out *Resp) error
		GetOrderById(ctx context.Context, in *GetOrderByIdReq, out *GetOrderByIdResp) error
		GetOrdersByUserId(ctx context.Context, in *GetOrdersByUserIdReq, out *GetOrdersByUserIdResp) error
		CancelOrder(ctx context.Context, in *CancelOrderReq, out *Resp) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) CreateOrder(ctx context.Context, in *CreateOrderReq, out *Resp) error {
	return h.OrderServiceHandler.CreateOrder(ctx, in, out)
}

func (h *orderServiceHandler) GetOrderById(ctx context.Context, in *GetOrderByIdReq, out *GetOrderByIdResp) error {
	return h.OrderServiceHandler.GetOrderById(ctx, in, out)
}

func (h *orderServiceHandler) GetOrdersByUserId(ctx context.Context, in *GetOrdersByUserIdReq, out *GetOrdersByUserIdResp) error {
	return h.OrderServiceHandler.GetOrdersByUserId(ctx, in, out)
}

func (h *orderServiceHandler) CancelOrder(ctx context.Context, in *CancelOrderReq, out *Resp) error {
	return h.OrderServiceHandler.CancelOrder(ctx, in, out)
}
