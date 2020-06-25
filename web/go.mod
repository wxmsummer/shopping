module shopping/web

go 1.13

replace github.com/wxmsummer/shopping/user/proto/user => ../user/proto/user

replace github.com/wxmsummer/shopping/product/proto/product => ../product/proto/product

replace github.com/wxmsummer/shopping/order/proto/order => ../order/proto/order

replace github.com/wxmsummer/shopping/comment/proto/comment => ../comment/proto/comment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/go-plugins/registry/consul/v2 v2.8.0
	github.com/rs/xid v1.2.1
	github.com/wxmsummer/shopping/comment/proto/comment v0.0.0-00010101000000-000000000000
	github.com/wxmsummer/shopping/order/proto/order v0.0.0-00010101000000-000000000000
	github.com/wxmsummer/shopping/product/proto/product v0.0.0-00010101000000-000000000000
	github.com/wxmsummer/shopping/user/proto/user v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
)
