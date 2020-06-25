package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shopping/web/handler"
)

func main() {

	// 注册consul地址
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	//// 使用selector来选择服务
	//for {
	//	getService, err := consulReg.GetService("go.micro.service.user")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// 轮询选择某个服务副本并调用
	//	node, err := selector.RoundRobin(getService)()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(node.Id, node.Address, node.Metadata)
	//	time.Sleep(time.Second * 1)
	//}

	//getService, err := consulReg.GetService("go.micro.service.user")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 随机选择某个服务副本并调用
	//node, err := selector.Random(getService)()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(node.Id, node.Address, node.Metadata)

	//mySelector := selector.NewSelector(
	//	selector.Registry(consulReg),
	//	selector.SetStrategy(selector.RoundRobin))

	// 把gin整合入micro框架中
	ginRouter := gin.Default()

	// create new web service
	// 创建web服务
	service := web.NewService(
		// 应该使用配置文件来设置name、address等
		web.Name("go.micro.service.web"),
		web.Version("latest"),
		web.Address(":8080"),
		web.Registry(consulReg),
		web.Handler(ginRouter),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	ginRouter.Static("/static", "./static")
	ginRouter.LoadHTMLGlob("views/**/*")

	ginRouter.GET("/", handler.GetIndex)

	// 设置url组
	userGroup := ginRouter.Group("/user")
	{
		userGroup.GET("/register", handler.GetUserRegister)
		userGroup.GET("/login", handler.GetUserLogin)
		userGroup.GET("/logout", handler.UserLogout)
		userGroup.GET("/getLevel", handler.GetUserLevel)
		userGroup.GET("/userCenter", handler.GetUserCenter)
		userGroup.GET("/myInfo", handler.GetMyInfo)
		userGroup.GET("/myOrder", handler.GetMyOrder)
		userGroup.GET("/myAddress", handler.GetMyAddress)
		userGroup.POST("/register", handler.PostUserRegister)
		userGroup.POST("/login", handler.PostUserLogin)
	}

	commentGroup := ginRouter.Group("/comment")
	{
		commentGroup.GET("/showComments/:productId", handler.GetComments)
		commentGroup.GET("/addComment", handler.GetAddComment)
		commentGroup.POST("/addComment", handler.PostAddComment)
	}

	productGroup := ginRouter.Group("/product")
	{
		productGroup.GET("/list", handler.GetProducts)
		productGroup.GET("/cart", handler.GetMyCart)
		productGroup.GET("/detail/:productId", handler.GetDetail)
		productGroup.GET("/search", handler.GetSearch)

		productGroup.POST("/search", handler.GetSearch)
	}

	orderGroup := ginRouter.Group("/order")
	{
		orderGroup.GET("/placeOrder", handler.GetPlaceOrder)
	}

	adminGroup := ginRouter.Group("/admin")
	{
		adminGroup.GET("/index", handler.GetAdminIndex)
		adminGroup.GET("/login", handler.GetAdminLogin)
		adminGroup.GET("/logout", handler.GetAdminIndex)
		adminGroup.GET("/addProduct", handler.GetAdminAddProduct)
		adminGroup.GET("/listProduct", handler.GetAdminListProduct)
		adminGroup.GET("/updateProduct", handler.GetAdminUpdateProduct)

		adminGroup.POST("/login", handler.PostAdminLogin)
		adminGroup.POST("/logout", handler.AdminLogout)
		adminGroup.POST("/addProduct", handler.PostAddProduct)
		adminGroup.POST("/updateProduct", handler.PostUpdateProduct)


	}

	// register product handler
	service.HandleFunc("/product/searchByID", handler.SearchByID)
	service.HandleFunc("/product/searchByName", handler.SearchByName)
	service.HandleFunc("/product/searchByClassify", handler.SearchByClassify)
	service.HandleFunc("/product/sortByPrice", handler.SortByPrice)
	service.HandleFunc("/product/sortBySalesVolume", handler.SortBySalesVolume)
	service.HandleFunc("/product/sortByCommentsNum", handler.SortByCommentsNum)
	service.HandleFunc("/product/delProduct", handler.DelProduct)

	// register order handler
	service.HandleFunc("/order/createOrder", handler.CreateOrder)
	service.HandleFunc("/order/getOrderById", handler.GetOrderById)
	service.HandleFunc("/order/getAllOrders", handler.GetAllOrders)
	service.HandleFunc("/order/cancelOrder", handler.CancelOrder)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
