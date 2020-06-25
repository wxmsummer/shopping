module shopping/order

go 1.13

replace shopping/order/proto/order => ./proto/order

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/jinzhu/gorm v1.9.14
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/go-plugins/registry/consul/v2 v2.8.0
	shopping/order/proto/order v0.0.0-00010101000000-000000000000
)
