package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/errors"
	"shopping/product/model"
	"shopping/product/repository"

	log "github.com/micro/go-micro/v2/logger"

	proto "shopping/product/proto/product"
)

type Product struct{Repo *repository.Product}

func (e *Product) SearchByID(ctx context.Context, req *proto.SearchByIDReq, rsp *proto.SearchResp) error {
	log.Info("Received Product.SearchById request")
	var products []*proto.Product
	err := e.Repo.Db.Where("id = ?" , req.Id).Find(&products).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("搜索结果共%v条" , len(products))
	rsp.Products = products

	return nil
}

func (e *Product) SearchByMethod(ctx context.Context, req *proto.SearchByMethodReq, rsp *proto.SearchResp) error {
	log.Info("Received Product.SearchByMethod request")
	var products []*proto.Product

	method := req.Method
	var err error
	switch method {
	case "Name":
		err = e.Repo.Db.Where("name like ?", "%" + req.SearchText + "%").Limit(10).Find(&products).Error
	case "Class":
		err = e.Repo.Db.Where("classify like ?", "%" + req.SearchText + "%").Limit(10).Find(&products).Error
	case "Tag":
		err = e.Repo.Db.Where("tag like ?", "%" + req.SearchText + "%").Limit(10).Find(&products).Error
	default:
		err = errors.New("502", "搜索方法不存在，仅支持 Name/Class/Tag", 502)
	}
	if err != nil {
		rsp.Code = 500
		rsp.Msg = "Db.Where err"
		return err
	}

	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("搜索结果共%v条" , len(products))
	rsp.Products = products

	return nil
}

func (e *Product) SortByNameAndMethod(ctx context.Context, req *proto.SortByNameAndMethodReq, rsp *proto.SortResp) error {
	log.Info("Received Product.SortByNameAndMethod request")
	var products []*proto.Product

	method := req.Method
	var err error
	tempRes := e.Repo.Db.Where("name like ?", "%" + req.Name + "%")
	switch method {
	case "Price":
		err = tempRes.Order("price desc").Limit(10).Find(&products).Error
	case "Sales":
		err = tempRes.Order("price desc").Limit(10).Find(&products).Error
	case "Comments":
		err = tempRes.Order("price desc").Limit(10).Find(&products).Error
	default:
		err = errors.New("502", "排序方法不存在，仅支持 Price/Sales/Comments", 502)
	}
	if err != nil {
		rsp.Code = 500
		rsp.Msg = "Db.Order err"
		return err
	}

	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("搜索结果共%v条，按评价数量降序排序" , len(products))
	rsp.Products = products

	return nil
}

func (e *Product) AddProduct(ctx context.Context, req *proto.AddProductReq, rsp *proto.Resp) error {
	log.Info("Received Product.AddProduct request")
	product := &model.Product{
		Name:        req.Product.Name,
		Classify:    req.Product.Classify,
		Tag:         req.Product.Tag,
		Price:       req.Product.Price,
		SalesVolume: 0,	// 销量默认为0
		CommentsNum: 0,	// 评论数默认为0
		Inventory:   req.Product.Inventory,
		Describe:    req.Product.Describe,
	}
	err := e.Repo.Create(product)
	if err != nil {
		log.Error("e.Repo.Create(proto) err=",err)
		return err
	}
	rsp.Code = 200
	rsp.Msg = "增加商品成功！"
	return nil
}

func (e *Product) UpdateProduct(ctx context.Context, req *proto.UpdateProductReq, rsp *proto.Resp) error {
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
	product.Describe = req.Product.Describe

	_, err = e.Repo.Update(product)
	if err != nil {
		return err
	}
	rsp.Code = 200
	rsp.Msg = "更新商品成功！"
	return nil
}

func (e *Product) DelProduct(ctx context.Context, req *proto.DelProductReq, rsp *proto.Resp) error {
	log.Info("Received Product.Call request")
	err := e.Repo.Delete(req.Id)
	if err != nil {
		return err
	}
	rsp.Code = 200
	rsp.Msg = fmt.Sprintf("删除编号为%d商品成功！", req.Id)
	return nil
}