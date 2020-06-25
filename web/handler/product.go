package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	proto "github.com/wxmsummer/shopping/product/proto/product"
	"net/http"
)

// 获取商品列表
func GetProducts(c *gin.Context) {
	c.HTML(http.StatusOK, "product_list.html", gin.H{
		"test_data": "test data show",
	})
}

// 获取购物车
func GetMyCart(c *gin.Context) {
	c.HTML(http.StatusOK, "product_cart.html", nil)
}

// 获取
func GetDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "product_detail.html", nil)
}

// 获取
func GetSearch(c *gin.Context) {
	c.HTML(http.StatusOK, "product_search.html", nil)
}

// 根据某方法搜索
func SearchByMethod(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	searchMethod := c.Query("SearchMethod")
	searchText := c.Query("SearchText")

	if searchText == "" {
		c.JSON(501, "searchText 为空字符串，请重新输入！")
	} else if searchMethod == "" {
		c.JSON(501, "searchMethod 为空字符串，请重新输入！")
	}

	client := proto.NewProductService("go.micro.service.product", server.Client())
	rsp, err := client.SearchByMethod(context.TODO(), &proto.SearchByMethodReq{
		Method: searchMethod, SearchText: searchText,
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

// 根据名字和排序方法排序
func SortByNameAndMethod(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	name := c.Query("ProductName")

	// 从前端请求获取sortMethod
	sortMethod := c.Query("SortMethod")

	client := proto.NewProductService("go.micro.service.product", server.Client())
	rsp, err := client.SortByNameAndMethod(context.TODO(), &proto.SortByNameAndMethodReq{
		Name: name, Method: sortMethod,
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}
	c.JSON(http.StatusOK, rsp)

}
