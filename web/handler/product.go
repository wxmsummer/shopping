package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	proto "github.com/wxmsummer/shopping/product/proto/product"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type ProductBundle struct {
	Count       int32
	LittlePrice int32
	Product     *proto.Product// 商品
}

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

// 获取购物车
func PostMyCart(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// 获取购物车所有id的cookie
	allProductId, err := c.Cookie("all_product_id")
	if err != nil {
		fmt.Println("c.Cookie() err=", err)
		return
	}
	
	var productIDs []string
	var productCounts map[string]int
	var totalCount int
	
	// 将cookie中的id和count拆分出来
	cookie_slice := strings.Split(allProductId, "AND")
	for _, productId:=range cookie_slice{
		productIDs = append(productIDs, productId)
		countStr, _ := c.Cookie(productId)
		count,_ :=strconv.Atoi(countStr)
		// 每个商品id对应一个count
		productCounts[productId] = count
		// 计算总数量
		totalCount += count
	}

	// 根据商品id封装req，从后端获取商品组
	client := proto.NewProductService("go.micro.service.product", server.Client())
	rsp, err := client.GetCartByProductIds(context.TODO(), &proto.GetCartReq{
		ProductIDs: productIDs,
	})

	// 获取总价
	var totalPrice float64
	var goods []ProductBundle
	var good ProductBundle
	for _, product := range rsp.Products{
		// 商品数量
		count := productCounts[product.ProductID]
		// 商品单价
		price, _ := strconv.ParseFloat(product.Price, 64)
		// 数量*单价
		littlePrice := float64(count) * price
		// 总价
		totalPrice += littlePrice
		// ProductBundle
		good = ProductBundle{
			Count:       int32(count),
			LittlePrice: int32(littlePrice),
			Product:     product,
		}
		goods = append(goods, good)
	}

	var data = map[string]interface{}{
		"goods":      goods,
		"TotalPrice": totalPrice,
		"TotalCount": totalCount,
	}

	t, err := template.ParseFiles("../template/cart.html")
	if err != nil {
		fmt.Println("template.ParseFiles err = ", err)
	}
	fmt.Println("goods:", goods)
	_ = t.Execute(c.Writer, data)
}

// 获取
func GetDetail(c *gin.Context) {
	server := grpc.NewService()
	server.Init()
	Id := c.Query("ProductId")
	fmt.Println("ProductId:", Id)

	if Id == "" {
		c.JSON(501, "ProductId 为空字符串，请重新输入！")
	}

	client := proto.NewProductService("go.micro.service.product", server.Client())
	rsp, err := client.SearchByProductID(context.TODO(), &proto.SearchByProductIDReq{
		ProductID: Id,
	})

	if err != nil {
		c.JSON(500, rsp)
		fmt.Println("ProductId: ", Id, "Not Found!")
		return
	}

	Products := rsp.Products
	fmt.Println("Product Len:", len(Products))

	if len(Products) != 1 {
		fmt.Println(" Len Error:", len(Products))
		return
	}
	Product := Products[0]

	var data = map[string]interface{}{
		"Product": Product,
	}
	fmt.Println("Product Name:", Product.Name)
	fmt.Println("Product:", Product)

	t, _ := template.ParseFiles("../template/detail.html")
	//t, _ :=template.ParseFiles("../web/views/product_search.html")
	fmt.Println(t.Name())
	t.Execute(c.Writer, data)
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
		c.JSON(501, "SearchText 为空字符串，请重新输入！")
	} else if searchMethod == "" {
		c.JSON(501, "SearchMethod 为空字符串，请重新输入！")
	}

	client := proto.NewProductService("go.micro.service.product", server.Client())
	rsp, err := client.SearchByMethod(context.TODO(), &proto.SearchByMethodReq{
		Method: searchMethod, SearchText: searchText,
	})
	if err != nil {
		c.JSON(500, rsp)
		return
	}

	goods := rsp.Products
	fmt.Println("goods len:", len(goods))
	fmt.Println(goods)

	data := map[string]interface{}{
		"goods": goods,
	}

	t, err := template.ParseFiles("../template/list.html")
	if err != nil {
		fmt.Println("template.ParseFiles err = ", err)
	}
	fmt.Println("goods:", goods)
	_ = t.Execute(c.Writer, data)

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
