package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/errors"
	"shopping/product/model"
	"shopping/product/repository"

	log "github.com/micro/go-micro/v2/logger"

	product "shopping/product/proto/product"
)

type Product struct{Repo *repository.Product}

func (e *Product) SearchByID(ctx context.Context, req *product.SearchByIDReq, rsp *product.SearchResp) error {
	log.Info("Received Product.SearchByName request")
	var products []*product.Product
	err := e.Repo.Db.Where("id = ?" , req.Id).Find(&products).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("搜索结果共%v条" , len(products))
	rsp.Products = products

	return nil
}

func (e *Product) SearchByName(ctx context.Context, req *product.SearchByNameReq, rsp *product.SearchResp) error {
	log.Info("Received Product.SearchByName request")
	var products []*product.Product
	err := e.Repo.Db.Where("name like ?", "%" + req.Name + "%").Limit(10).Find(&products).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("搜索结果共%v条" , len(products))
	rsp.Products = products

	return nil
}

func (e *Product) SearchByClassify(ctx context.Context, req *product.SearchByClassifyReq, rsp *product.SearchResp) error {
	log.Info("Received Product.SearchByClassify request")
	var products []*product.Product
	err := e.Repo.Db.Where("classify like ?", "%" + req.Classify + "%").Limit(10).Find(&products).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("搜索结果共%v条" , len(products))
	rsp.Products = products

	return nil
}

func (e *Product) SearchByTag(ctx context.Context, req *product.SearchByTagReq, rsp *product.SearchResp) error {
	log.Info("Received Product.SearchByClassify request")
	var products []*product.Product
	err := e.Repo.Db.Where("tag like ?", "%" + req.Tag + "%").Limit(10).Find(&products).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("搜索结果共%v条" , len(products))
	rsp.Products = products

	return nil
}

func (e *Product) SortByPrice(ctx context.Context, req *product.Request, rsp *product.Response) error {
	log.Info("Received Product.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Product) SortBySalesVolume(ctx context.Context, req *product.Request, rsp *product.Response) error {
	log.Info("Received Product.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Product) SortByCommentsNum(ctx context.Context, req *product.Request, rsp *product.Response) error {
	log.Info("Received Product.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Product) AddProduct(ctx context.Context, req *product.AddProductReq, rsp *product.Resp) error {
	log.Info("Received Product.AddProduct request")
	product := &model.Product{
		Name:         req.Product.Name,
		Classify:     req.Product.Classify,
		Tag:          req.Product.Tag,
		Price:        req.Product.Price,
		SalesVolume:  0,	// 销量默认为0
		CommentsNum:  0,	// 评论数默认为0
		Inventory:    req.Product.Inventory,
		Introduction: req.Product.Introduction,
	}
	err := e.Repo.Create(product)
	if err != nil {
		log.Error("e.Repo.Create(product) err=",err)
		return err
	}
	rsp.Code = 200
	rsp.Msg = "增加商品成功！"
	return nil
}

func (e *Product) UpdateProduct(ctx context.Context, req *product.UpdateProductReq, rsp *product.Resp) error {
	log.Info("Received Product.UpdateProduct request")

	product, err := e.Repo.FindByID(req.Product.Id)
	if err != nil {
		return errors.Unauthorized("go.micro.srv.Product.UpdateProduct", "该商品不存在")
	}
	product.Name= req.Product.Name
	product.Classify= req.Product.Classify
	product.Tag= req.Product.Tag
	product.Price= req.Product.Price
	product.SalesVolume= req.Product.SalesVolume
	product.CommentsNum= req.Product.CommentsNum
	product.Inventory= req.Product.Inventory
	product.Introduction= req.Product.Introduction

	_, err = e.Repo.Update(product)
	if err != nil {
		return err
	}
	rsp.Code = 200
	rsp.Msg = "更新商品成功！"
	return nil
}

func (e *Product) DelProduct(ctx context.Context, req *product.Request, rsp *product.Response) error {
	log.Info("Received Product.Call request")
	return nil
}