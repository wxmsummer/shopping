// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/product/product.proto

package go_micro_service_product

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

// Client API for ProductService service

type ProductService interface {
	SearchByID(ctx context.Context, in *SearchByIDReq, opts ...client.CallOption) (*SearchResp, error)
	SearchByName(ctx context.Context, in *SearchByNameReq, opts ...client.CallOption) (*SearchResp, error)
	SearchByClassify(ctx context.Context, in *SearchByClassifyReq, opts ...client.CallOption) (*SearchResp, error)
	SearchByTag(ctx context.Context, in *SearchByTagReq, opts ...client.CallOption) (*SearchResp, error)
	SortByPrice(ctx context.Context, in *SortByPriceReq, opts ...client.CallOption) (*SortResp, error)
	SortBySalesVolume(ctx context.Context, in *SortBySalesVolumeReq, opts ...client.CallOption) (*SortResp, error)
	SortByCommentsNum(ctx context.Context, in *SortByCommentsNumReq, opts ...client.CallOption) (*SortResp, error)
	AddProduct(ctx context.Context, in *AddProductReq, opts ...client.CallOption) (*Resp, error)
	UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...client.CallOption) (*Resp, error)
	DelProduct(ctx context.Context, in *DelProductReq, opts ...client.CallOption) (*Resp, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.service.product"
	}
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) SearchByID(ctx context.Context, in *SearchByIDReq, opts ...client.CallOption) (*SearchResp, error) {
	req := c.c.NewRequest(c.name, "ProductService.SearchByID", in)
	out := new(SearchResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) SearchByName(ctx context.Context, in *SearchByNameReq, opts ...client.CallOption) (*SearchResp, error) {
	req := c.c.NewRequest(c.name, "ProductService.SearchByName", in)
	out := new(SearchResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) SearchByClassify(ctx context.Context, in *SearchByClassifyReq, opts ...client.CallOption) (*SearchResp, error) {
	req := c.c.NewRequest(c.name, "ProductService.SearchByClassify", in)
	out := new(SearchResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) SearchByTag(ctx context.Context, in *SearchByTagReq, opts ...client.CallOption) (*SearchResp, error) {
	req := c.c.NewRequest(c.name, "ProductService.SearchByTag", in)
	out := new(SearchResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) SortByPrice(ctx context.Context, in *SortByPriceReq, opts ...client.CallOption) (*SortResp, error) {
	req := c.c.NewRequest(c.name, "ProductService.SortByPrice", in)
	out := new(SortResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) SortBySalesVolume(ctx context.Context, in *SortBySalesVolumeReq, opts ...client.CallOption) (*SortResp, error) {
	req := c.c.NewRequest(c.name, "ProductService.SortBySalesVolume", in)
	out := new(SortResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) SortByCommentsNum(ctx context.Context, in *SortByCommentsNumReq, opts ...client.CallOption) (*SortResp, error) {
	req := c.c.NewRequest(c.name, "ProductService.SortByCommentsNum", in)
	out := new(SortResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) AddProduct(ctx context.Context, in *AddProductReq, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "ProductService.AddProduct", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "ProductService.UpdateProduct", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) DelProduct(ctx context.Context, in *DelProductReq, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "ProductService.DelProduct", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceHandler interface {
	SearchByID(context.Context, *SearchByIDReq, *SearchResp) error
	SearchByName(context.Context, *SearchByNameReq, *SearchResp) error
	SearchByClassify(context.Context, *SearchByClassifyReq, *SearchResp) error
	SearchByTag(context.Context, *SearchByTagReq, *SearchResp) error
	SortByPrice(context.Context, *SortByPriceReq, *SortResp) error
	SortBySalesVolume(context.Context, *SortBySalesVolumeReq, *SortResp) error
	SortByCommentsNum(context.Context, *SortByCommentsNumReq, *SortResp) error
	AddProduct(context.Context, *AddProductReq, *Resp) error
	UpdateProduct(context.Context, *UpdateProductReq, *Resp) error
	DelProduct(context.Context, *DelProductReq, *Resp) error
}

func RegisterProductServiceHandler(s server.Server, hdlr ProductServiceHandler, opts ...server.HandlerOption) error {
	type productService interface {
		SearchByID(ctx context.Context, in *SearchByIDReq, out *SearchResp) error
		SearchByName(ctx context.Context, in *SearchByNameReq, out *SearchResp) error
		SearchByClassify(ctx context.Context, in *SearchByClassifyReq, out *SearchResp) error
		SearchByTag(ctx context.Context, in *SearchByTagReq, out *SearchResp) error
		SortByPrice(ctx context.Context, in *SortByPriceReq, out *SortResp) error
		SortBySalesVolume(ctx context.Context, in *SortBySalesVolumeReq, out *SortResp) error
		SortByCommentsNum(ctx context.Context, in *SortByCommentsNumReq, out *SortResp) error
		AddProduct(ctx context.Context, in *AddProductReq, out *Resp) error
		UpdateProduct(ctx context.Context, in *UpdateProductReq, out *Resp) error
		DelProduct(ctx context.Context, in *DelProductReq, out *Resp) error
	}
	type ProductService struct {
		productService
	}
	h := &productServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductService{h}, opts...))
}

type productServiceHandler struct {
	ProductServiceHandler
}

func (h *productServiceHandler) SearchByID(ctx context.Context, in *SearchByIDReq, out *SearchResp) error {
	return h.ProductServiceHandler.SearchByID(ctx, in, out)
}

func (h *productServiceHandler) SearchByName(ctx context.Context, in *SearchByNameReq, out *SearchResp) error {
	return h.ProductServiceHandler.SearchByName(ctx, in, out)
}

func (h *productServiceHandler) SearchByClassify(ctx context.Context, in *SearchByClassifyReq, out *SearchResp) error {
	return h.ProductServiceHandler.SearchByClassify(ctx, in, out)
}

func (h *productServiceHandler) SearchByTag(ctx context.Context, in *SearchByTagReq, out *SearchResp) error {
	return h.ProductServiceHandler.SearchByTag(ctx, in, out)
}

func (h *productServiceHandler) SortByPrice(ctx context.Context, in *SortByPriceReq, out *SortResp) error {
	return h.ProductServiceHandler.SortByPrice(ctx, in, out)
}

func (h *productServiceHandler) SortBySalesVolume(ctx context.Context, in *SortBySalesVolumeReq, out *SortResp) error {
	return h.ProductServiceHandler.SortBySalesVolume(ctx, in, out)
}

func (h *productServiceHandler) SortByCommentsNum(ctx context.Context, in *SortByCommentsNumReq, out *SortResp) error {
	return h.ProductServiceHandler.SortByCommentsNum(ctx, in, out)
}

func (h *productServiceHandler) AddProduct(ctx context.Context, in *AddProductReq, out *Resp) error {
	return h.ProductServiceHandler.AddProduct(ctx, in, out)
}

func (h *productServiceHandler) UpdateProduct(ctx context.Context, in *UpdateProductReq, out *Resp) error {
	return h.ProductServiceHandler.UpdateProduct(ctx, in, out)
}

func (h *productServiceHandler) DelProduct(ctx context.Context, in *DelProductReq, out *Resp) error {
	return h.ProductServiceHandler.DelProduct(ctx, in, out)
}
