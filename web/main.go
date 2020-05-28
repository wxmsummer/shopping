package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	"net/http"
	"shopping/web/handler"
)

func main() {

	consulReg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"127.0.0.1:8500",
		}
	})

	// create new web service
	service := web.NewService(
		web.Name("go.micro.service.web"),
		web.Version("latest"),
		web.Address(":8080"),
		web.Registry(consulReg),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register user handler
	service.HandleFunc("/user/register", handler.Register)
	service.HandleFunc("/user/login", handler.Login)
	service.HandleFunc("/user/logout", handler.Logout)
	service.HandleFunc("/user/getLevel", handler.GetLevel)

	// register comment handler
	service.HandleFunc("/comment/addComment", handler.AddComment)
	service.HandleFunc("/comment/getComments", handler.GetComments)

	// register product handler
	service.HandleFunc("/product/searchByID", handler.SearchByID)
	service.HandleFunc("/product/searchByName", handler.SearchByName)
	service.HandleFunc("/product/searchByClassify", handler.SearchByClassify)
	service.HandleFunc("/product/sortByPrice", handler.SortByPrice)
	service.HandleFunc("/product/sortBySalesVolume", handler.SortBySalesVolume)
	service.HandleFunc("/product/sortByCommentsNum", handler.SortByCommentsNum)
	service.HandleFunc("/product/addProduct", handler.AddProduct)
	service.HandleFunc("/product/updateProduct", handler.UpdateProduct)
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
