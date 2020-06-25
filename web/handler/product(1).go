package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	"html/template"
)

type Product struct {
	Name        string // 商品名
	Classify    string // 分类
	Tag         string // 标签
	Price       string // 价格
	SalesVolume int32  // 销量
	CommentsNum int32  // 评价数量
	Inventory   int32  // 库存
	Describe    string // 商品描述
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

    fmt.Println("searchText:",searchText)
    fmt.Println("searchMethod:",searchMethod)
    /*
	client := proto.NewProductService("go.micro.service.product", server.Client())
	 client.SearchByMethod(context.TODO(), &proto.SearchByMethodReq{
		Method: searchMethod, SearchText: searchText,
	})
    */

    var goods = make([]Product, 0)
                good1:=Product{
                    Name:"test1",
                	Classify:"手机",
                	Tag :"好点",
                	Price:"123123",
                	SalesVolume:333,
                	CommentsNum:444,
                	Inventory :22222,
                	Describe:"真的牛逼",
                }
                fmt.Println("good1:",good1.Name)


                 good2:=Product{
                        Name:"test2",
                    	Classify:"电脑",
                    	Tag :"垃圾",
                    	Price:"333333",
                    	SalesVolume:111,
                    	CommentsNum:111,
                    	Inventory :333333,
                    	Describe:"真的垃圾",
                    }
                fmt.Println("good2:",good2.Name)

                goods = append(goods,good1)
                goods = append(goods,good2)

                fmt.Println("goods len:",len(goods))

            	var data = map[string]interface{}{
                			"goods":     goods,
                }
                fmt.Println("goods:",goods)

        t, _ :=template.ParseFiles("../template/layout.html")
        fmt.Println(t.Name())
        t.Execute(c.Writer,data)
}
